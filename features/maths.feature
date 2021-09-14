Feature: I should be able to do simple maths

  Scenario: I can square numbers
    Given an initiated squares stream
    And no cached numbers
    When I add 1 to the stream
     And I add 2 to the stream
     And I add 3 to the stream
     And I add 4 to the stream
     And I add 5 to the stream
    Then the sum of cached squares should be 55
