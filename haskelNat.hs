factorial n = if n > 1
              then n * factorial (n - 1)
              else 1
main = do
    print ("Input A Number To Get It's Factorial")
    name <- readLn
    print (factorial name)
