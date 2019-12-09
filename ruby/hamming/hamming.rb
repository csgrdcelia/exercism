class Hamming
  def self.compute(first_string, second_string)
    raise ArgumentError if first_string.length != second_string.length
    result = 0
    for position in 0...first_string.length
      result += 1 if first_string[position] != second_string[position]
    end 
    result
  end
end
