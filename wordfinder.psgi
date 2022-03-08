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

# HTML-free errors.
sub simple_error {
     my ($message, $code) = @_;
     status $code || 404;
     return $message;
 }

get '/ping'              => sub { 'ok' };
get '/wordfinder/:input' => sub { &wordfinder };
any qr{.*}               => sub { simple_error("Not Found", 404) };

#
# wordfinder() finds all words built from a list of letters.
#
# The words can only contain letters from the list.
# Not all letters need to be used.
# Letters can appear in the words, at most, as many times as they appear in the list.
#
# Warning: Nested loops ahead!
my @WORDS;
sub wordfinder {
    my $input = lc route_parameters->get('input');
    $input =~ s{[^a-z]}{}gxms;    # Ignore anything outside of a-z. TODO: return error?
    my $input_len = length $input;
    return simple_error( "Bad Request", 400 ) unless $input_len;
    return simple_error( "Bad Request", 400 ) if $input_len > 26;
    load_words() unless @WORDS;

    # Keep track of the how many times each letter appears in the input list.
    my %input_counts;
    foreach ( split //, $input) {
        $input_counts{$_} ++;
    }
    
    my $pattern = qr{\A [\Q$input\E]+ \z}xms; # Could use keys %input_counts here I guess.

    my @matches;
    # This is a bottleneck.
    # Maybe collapse the words into a space-separated string and make one regex pass through that??
    foreach (@WORDS) {
        next if length $_ > $input_len; # Fast and cheap.
        push @matches, $_ if $_ =~ m{$pattern};
    }
    # Matches includes overmatching. ban => banana
    # Process the matches again to filter out those cases.
    # This is death by nested loops, but the sample set should hopefully be small enough.
    my @filtered_matches;
    MATCH: foreach my $match ( @matches ) {
        my %match_counts;
        foreach ( split //, $match) {
            $match_counts{$_} ++;
            next MATCH if $match_counts{$_} > $input_counts{$_};
        }
        push @filtered_matches, $match;
    }
    send_as 'JSON' => \@filtered_matches;
}

# Read all the words from the dictionary file, filter to a-z, convert to
# lowercase, and remove dups.
sub load_words {
    open my $FH, "<", "/usr/share/dict/words" or die $!;
    @WORDS = <$FH>;
    close $FH or die $!;
    chomp @WORDS;
    my %WORDS = map { lc $_ => 1 } @WORDS;

    @WORDS = sort grep { m{ \A [a-z]+ \z }xms } keys %WORDS;
}

# Must be the last command in file.
__PACKAGE__->to_app;
