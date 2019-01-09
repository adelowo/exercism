class Acronym
  def self.abbreviate(word)
    abbreviation = ''

    word.split(' ').each do |w|
      w.upcase!

      add = false

      splitted = w.split('-')

      if splitted.length.eql? 2
        abbreviation += "#{splitted[0][0]}#{splitted[1][0]}"
        next
      end

      next unless w.split('').each do |char|
        add = true if ('A'..'Z').cover? char
      end

      abbreviation += w[0] if add == true
    end

    abbreviation
  end
end
