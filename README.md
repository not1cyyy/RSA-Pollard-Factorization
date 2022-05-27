# What is Pollard P-1 Algorithm ? 
According to wikipedia.com, Pollard's p − 1 algorithm is a number theoretic integer factorization algorithm, invented by John Pollard in 1974. It is a special-purpose algorithm, meaning that it is only suitable for integers with specific types of factors; it is the simplest example of an algebraic-group factorisation algorithm. 

# Why Pollard P-1 Algorithm ?
Sometimes the integers aren't that big to factorize using factordb.com, so it's relatively easier to use a script that helps factorizing these integers 

# How does Pollard P-1 Algorithm work ? 

Given a number n.
Initialize a = 2, i = 2
Until a factor is returned do
a <- (a^i) mod n
d <- GCD(a-1, n)
if 1 < d < n then
    return d
else
    i <- i+1
Other factor, d' <- n/d
If d' is not prime
    n <- d'
    goto 1
else
    d and d' are two prime factors.


# FAQ
If you have any questions don't hesitate to reach out to me ! 