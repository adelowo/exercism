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
    scores.max 3
  end

  def report
    "Your latest score was #{latest}. #{report_details}"
  end

  private

  def report_details
    latest_score = latest

    if personal_best.eql? latest_score
      return "That's your personal best!"
    else

      short = personal_top.first - latest_score
      "That's #{short} short of your personal best!" if short > 0
    end
  end
end
