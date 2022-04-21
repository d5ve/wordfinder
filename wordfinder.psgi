#!/usr/bin/env perl

=head1 NAME

wordfinder.psgi - wordfinder HTTP API

=head1 SYNOPSIS

    plackup wordfinder.psgi --listen :80

=head1 DESCRIPTION

Simple perl/dancer2 JSON API to find words matching a pattern.

When supplied a list of letters, return a list of all the words that can be
built from just those letters. The input letters can include repeats. Not all
input letters need to be used. Each letter in the input list can appear at most
once in the output words.

=head1 LICENSE

Wordfinder is free software; you can redistribute it and/or modify it under the
same terms as perl itself.

=cut

use Dancer2 appname => 'wordfinder';

our $VERSION = v0.4;

# load_words() hits disk.
my @WORDS             = load_words();
my $sane_input_length = 26; # What is a sane maximum here ??

# Routes
get '/ping'               => sub { 'ok' };
get '/wordfinder/:input'  => sub { &wordfinder };   # Loop through the wordlist checking each word.
get '/wordfinder4/:input' => sub { &wordfinder4 };  # Treat the wordlist as a single string and run a single regex. Faster. More mem??
any qr{.*}                => sub { status 404; return "Not Found"; };

#
# wordfinder() finds all words built from a list of letters.
#
# The words can only contain letters from the list.
# Not all letters need to be used.
# Letters can appear in the words, at most, as many times as they appear in the list.
#
# There is likely a simple method, that I slept through in university, for
# speeding this all up. It seems like looping through the whole list each time
# can't be necessary. I couldn't think of a tree structure that would be good
# for this "any order is fine" search pattern. Prefixes don't help. In my
# benchmarking, my laptop was looping through the words and running the regex
# about 4 million words per second, so maybe this is good enough!
#
# Warning: Nested loops ahead! (BRRRRRRR)
# wordfinder() currently makes one full pass through @WORDS, running a regex
# match on any words short enough to be fully matched by the input.
# Then each match is looped through per character to check that it uses a
# correct number of each character.
sub wordfinder {
    my $input = lc route_parameters->get('input');
    $input =~ s{[^a-z]}{}gxms;    # Ignore anything outside of a-z. TODO: return error?
    my $input_len = length $input;
    if ( !$input_len ) {
        status 404;
        return "Bad Request";
    }
    if ( $input_len > $sane_input_length ) {
        status 404;
        return "Bad Request";
    }

    # Keep track of the how many times each letter appears in the input list.
    my %input_counts;
    foreach ( split //, $input ) {
        $input_counts{$_}++;
    }

    # Bail at the first sign of any non-input char.
    my $pattern = qr{ [^\Q$input\E] }xms; # Could use keys %input_counts here I guess.
    my @matches;
    # This is a bottleneck.
    # Maybe collapse the words into a space-separated string and make one regex pass through that??
    foreach (@WORDS) {
        next if length $_ > $input_len;    # Fast and cheap.
        push @matches, $_ unless $_ =~ m{$pattern};
    }

    # Matches is now words that contain only the distinct letters in the input,
    # but not necessarily the correct number of each letter. banaaa => banana
    # Process the matches again to filter out those cases.
    # This is death by nested loops, but the sample set should hopefully be small enough.
    my @filtered_matches;
    MATCH: foreach my $match (@matches) {
        my %match_counts;
        foreach ( split //, $match ) {
            $match_counts{$_}++;
            next MATCH if $match_counts{$_} > $input_counts{$_};
        }
        push @filtered_matches, $match;
    }
    send_as 'JSON' => \@filtered_matches;
}

#
# New algorithm that puts the whole dictionary into a single string, then runs
# a single pattern match along the whole thing. This takes half the time per
# search as the original loop-style algorithm in wordfinder(), at the expense of
# RAM.
# See wordfinder() for more docs.
my $single_line_words;
sub wordfinder4 {
    my $input = lc route_parameters->get('input');
    $input =~ s{[^a-z]}{}gxms;    # Ignore anything outside of a-z. TODO: return error?
    my $input_len = length $input;
    if ( !$input_len ) {
        status 404;
        return "Bad Request";
    }
    if ( $input_len > $sane_input_length ) {
        status 404;
        return "Bad Request";
    }

    $single_line_words //= join(' ', @WORDS); # Lazily load a single string holding all the words.

    # Keep track of the how many times each letter appears in the input list.
    my %input_counts;
    foreach ( split //, $input ) {
        $input_counts{$_}++;
    }

    my $pattern = qr{ \b ([\Q$input\E]+) \b }xms; # Could use keys %input_counts here I guess.
    # Get a list of all matches in the dictionary string that are no longer than the input pattern.
    my @matches = grep { length $_ <= $input_len } $single_line_words =~ m{$pattern}g;

    # Matches is now words that contain only the distinct letters in the input,
    # but not necessarily the correct number of each letter. banaaa => banana
    # Process the matches again to filter out those cases.
    # This is death by nested loops, but the sample set should hopefully be small enough.
    my @filtered_matches;
    MATCH: foreach my $match (@matches) {
        my %match_counts;
        foreach ( split //, $match ) {
            $match_counts{$_}++;
            next MATCH if $match_counts{$_} > $input_counts{$_};
        }
        push @filtered_matches, $match;
    }
    send_as 'JSON' => \@filtered_matches;
}

# Read all the words from the dictionary file, filter to a-z, convert to
# lowercase, remove dups and the erroneous single-letter words.
sub load_words {
    open my $FH, "<", "/usr/share/dict/words" or die $!;
    my @words = <$FH>;
    close $FH or die $!;
    chomp @words;
    @words = grep { m{ \A [a-z]+ \z }ixms } @words;    # Filter out words with non-a-z.
    my %wordsmap = map { lc $_ => 1 } @words;         # Hashify to clobber dups.

    # This is the only English-specific bit of code (barring the a-z criteria).
    # Remove the erroneous single-letter words from hash.
    delete @wordsmap{ ( "b" .. "h", "j" .. "n", "p" .. "z" ) };

    return sort keys %wordsmap;    # This gets the JSON responses sorted "for free".
}

# Must be the last command in file.
__PACKAGE__->to_app;
