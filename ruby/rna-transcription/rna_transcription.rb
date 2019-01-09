class Complement
  STRANDS = { 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U' }.freeze

  def self.of_dna(nucleotide)
    rna = ''

    nucleotide.split('').each do |val|
      return '' if STRANDS[val].nil?

      rna += STRANDS[val]
    end
    rna
  end
end

module BookKeeping
  VERSION = 4
end
