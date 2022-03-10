#!/usr/bin/env perl

use strict;
use warnings;

use Data::Dumper; $Data::Dumper::Sortkeys = 1;
use Benchmark qw(:all);

# Test any speedup from a new regex.
my @WORDS            = load_words();
my $input            = "eariotnslc";    # Most common ten letters in English.
my $input_length     = length $input;
my $positive_pattern = qr{\A [\Q$input\E]+ \z}xms;
my $negative_pattern = qr{ [^\Q$input\E] }xms;
my %input            = map { $_ => 1 } split //, $input;
my $inverse          = join '', ( grep { !$input{$_} } ( 'a' .. 'z' ) );
my $inverse_pattern  = qr{ [\Q$inverse\E] }xms;

# Make all the
my $single_string  = join( " ", @WORDS );
my $single_pattern = qr{ \b ([\Q$input\E]+) \b }xms;

print '@WORDS has ' . ( scalar @WORDS ) . " words\n";

my $pmatches;
my $nmatches;
my $imatches;
my $smatches;

cmpthese(
    -2,
    {
        POSITIVE => sub {
            my $matches;
            foreach (@WORDS) {
                # next if length $_ > $input_length;
                $matches++ if $_ =~ m{$positive_pattern};
            }
            $pmatches->{$matches} ++;
        },
        NEGATIVE => sub {
            my $matches;
            foreach (@WORDS) {
                # next if length $_ > $input_length;
                $matches++ unless $_ =~ m{$negative_pattern};
            }
            $nmatches->{$matches} ++;
        },
        INVERSE => sub {
            my $matches;
            foreach (@WORDS) {
                # next if length $_ > $input_length;
                $matches++ unless $_ =~ m{$inverse_pattern};
            }
            $imatches->{$matches} ++;
        },
        SINGLE => sub {
            my $matches;
            $matches++ while ( $single_string =~ m{$single_pattern}g );
            $smatches->{$matches} ++;
        }
    }
);

print Dumper({
    POSITIVE => $pmatches,
    NEGATIVE => $nmatches,
    INVERSE  => $imatches,
    SINGLE   => $smatches,
});
exit;

# Read all the words from the dictionary file, filter to a-z, convert to
# lowercase, remove dups and the erroneous single-letter words.
sub load_words {
    open my $FH, "<", "/usr/share/dict/words" or die $!;
    my @words = <$FH>;
    close $FH or die $!;
    chomp @words;
    @words =
      grep { m{ \A [a-z]+ \z }xms } @words;    # Filter out words with non-a-z.
    my %wordsmap = map { lc $_ => 1 } @words;  # Hashify to clobber dups.

    # This is the only English-specific bit of code (barring the a-z criteria).
    # Remove the erroneous single-letter words from hash.
    delete @wordsmap{ ( "b" .. "h", "j" .. "n", "p" .. "z" ) };

    return
      sort keys %wordsmap;    # This gets the JSON responses sorted "for free".
}
