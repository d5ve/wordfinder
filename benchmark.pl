#!/usr/bin/env perl

use strict;
use warnings;

use Benchmark qw(:all);

# Test any speedup from a new regex.
my @WORDS = load_words();
my $input = "eariotnslc"; # Most common ten letters in English.
my $positive_pattern = qr{\A [\Q$input\E]+ \z}xms;
my $negative_pattern = qr{ [^\Q$input\E] }xms;

cmpthese(
    -2,
    {
        POSITIVE => sub {
            my @matches;
            foreach (@WORDS) {
                push @matches, $_ if $_ =~ m{$positive_pattern};
            }
            # print 'P' . scalar @matches, "\n";
        },
        NEGATIVE => sub {
            my @matches;
            foreach (@WORDS) {
                push @matches, $_ unless $_ =~ m{$negative_pattern};
            }
            # print 'N' . scalar @matches, "\n";
        },
    }
);

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
