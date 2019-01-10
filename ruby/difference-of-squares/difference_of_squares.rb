class Squares
  attr_reader :value

  def initialize(value)
    @value = value
  end

  def square_of_sum
    return 1 if @value.eql? 1

    total = 0
    (1..@value).each do |i|
      total += i
    end

    total**2
  end

  def sum_of_squares
    return 1 if @value.eql? 1

    (1..@value).reduce do |total, num|
      total + (num**2)
    end
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
