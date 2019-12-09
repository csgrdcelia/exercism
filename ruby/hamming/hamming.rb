class Hamming
  def self.compute(first_string, second_string)
    raise ArgumentError if first_string.length != second_string.length
    0.upto(first_string.length).count { |i| first_string[i] != second_string[i] }
  end
end
