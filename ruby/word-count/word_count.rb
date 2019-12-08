class Phrase

    def initialize(phrase)
        @phrase = phrase.split(/[^[[:word:]]']+/).reject { |c| c.empty? }
    end

    def word_count
        words = Hash.new(0)
        @phrase.each { |word| words[word.downcase] += 1 }
        words
    end

end

