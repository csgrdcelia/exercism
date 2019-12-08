class Phrase

    def initialize(phrase)
        @phrase = phrase.downcase.split(/[\s,.!:?@$%^&]/).reject { |c| c.empty? }
    end

    def word_count
        @phrase.each_with_object(Hash.new(0)) {|word, counts| counts[word] += 1 }
    end

end

