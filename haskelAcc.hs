factorial n = let loop acc n' = if n' > 1
                                then loop (acc * n') (n' - 1)
                                else acc
               in loop 1 n
main = do
    print ("Input A Number To Get It's Factorial")
    name <- readLn
    print (factorial name)
