# Project




# WHat everything does




# Steps to take to do each part




# architecture


# Tests

- Unit Testing


Using TDD and Gos Testing package

Test Driven Development
1. Write test code before application code

Operates in the Red --> Green --> Refactor Loop

🔴 RED: Write a test that fails (or won't compile)
          │
          ▼
    🟢 GREEN: Write the minimum production code to make it pass
          │
          ▼
    🛠️ REFACTOR: Clean up code & remove duplication while staying green

    🔴 Red: Write a test targeting a feature that doesn't exist yet. Running go test should fail (either because it doesn't compile or an assertion fails).
    
    🟢 Green: Write the absolute simplest code required to make the test pass. Do not worry about clean code here; just fix the failure.
    
    🛠️ Refactor: Clean up both your application code and your test structure. Optimize logic, improve naming, and eliminate duplication while ensuring go test remains green



## Test Package

*Use Case*: Automate testcases of packages

*func signature*: 

```go
func TestXxx(*testing.T)

// Where func name starts with cap, and its input is the function you are testing so we put a pointer to that testing function
```

Withn These functions, we use

T.Error, T.Fail or related methods to singal failure

the file named "_test.go". will be excluded from regular package builds but will be included when the "go test" command is run

Test file can be in the sme pckage as the one being tested, or in another file with suffic *"_test.go"





