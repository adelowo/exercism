//Solution goes in Sources

class Year {
    
    private var yearToCheck : Int
    
    public var isLeapYear: Bool {
        get {
          return (yearToCheck % 4 == 0 ||
            (yearToCheck % 100 == 0 && yearToCheck % 400 == 0))
        }
    }
    
    init(calendarYear : Int) {
        self.yearToCheck = calendarYear
    }
}
