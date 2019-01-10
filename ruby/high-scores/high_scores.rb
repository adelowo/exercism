class HighScores
  attr_reader :scores

  def initialize(all_scores)
    @scores = all_scores
  end

  def latest
    scores.last
  end

  def personal_best
    scores.max
  end

  def personal_top
    scores.sort_by(&:-@)[0..2]
  end

  def report
    latest_score = latest
    return "Your latest score was #{latest_score}. That's your personal best!" if personal_best.eql? latest_score

    short = personal_top.first - latest_score
    return "Your latest score was #{latest_score}. That's #{short} short of your personal best!" if short > 0
  end

end
