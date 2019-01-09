class Hamming
  def self.compute(first, second)
    raise ArgumentError unless first.length.equal?(second.length)

    matches = 0

    second = second.split('')

    first.each_char do |char|
      matches += 1 unless char.eql?(second.shift)
    end

    matches
  end
end

module BookKeeping
  VERSION = 3
end
