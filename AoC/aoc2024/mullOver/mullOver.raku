my $fh = open "./input.txt", :r;
my $corruptedMemory = $fh.slurp;
$fh.close;

sub processMemory($mem) {
  my $sum = 0
  for $mem.match(/ 'mul(' (\d{1..3}) ',' (\d{1..3}) ')' /, :g) -> $n, $m{
    $sum += 1
  }
}