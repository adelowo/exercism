#!/bin/bash

## Bash script to run all the tests at a go

for i in $(ls); do

  if [[ $i == "go" ]]; then
    echo "Running the Go vet command"

    go vet ./...

    echo "Vet command done"

    echo "Running the testsuite for the Go track"

    go test ./...
    continue
  fi

  if [[ $i == "php" ]]; then
    echo ""
    echo "Running testsuite in PHP"
    ### Why write a `phpunit.xml` file ?
    ## when you can do that stuff in the shell asap
    find $i -type f -name "*test.php" | xargs phpunit --colors
    continue
  fi

  if [[ $i == "ruby" ]]; then
    echo ""
    echo "Running the testsuite for the Ruby track"

    find $i -type f -name "*_test.rb" | xargs ruby
    continue
  fi
done
