#!/usr/bin/env python
#coding: utf-8
######
# http://projecteuler.net/problems
######

import math
import decimal
import string
import itertools
import re
from fractions import Fraction # problem #243


#################################################
# Helper functions

def sieve_of_eratosthenes(limit):
    primes = range(2,limit)
    #i = 0
    #print "primes array built"
    #while i <= limit**0.5:
    #    print primes[i]
    #    for j in primes[(i+1):]:
    #        if not j%primes[i]:
    #            primes.remove(j)
    #    i+=1
    #    print i
    for i in range(2, int(limit**0.5)):
        primes = filter(lambda x: x == i or x % i, primes)
    return primes

# taken from http://stackoverflow.com/a/3035188/
def sieve_of_eratosthenes_2(limit):
    primes = [True] * limit                                     # build a list of length limit
    for i in xrange(3,int(limit**0.5)+1, 2):                    # for all odd numbers up to the square root of limit
        if primes[i]:                                           # only check numbers that haven't already been 'removed'
            primes[i*i::2*i]=[False]*((limit-i*i-1)/(2*i)+1)    # starting with prime 'i', remove i**2 and (i**2)+2n
    return [2] + [i for i in xrange(3,limit,2) if primes[i]]    # return the remaining list

# generate permutations of rotations of a string (e.g. abcde: (abcde, bcdea, cdeab, deabc, eabcd))
def permurotate(s):
    permutations = []
    i = 0
    while i < len(s):
        permutations.append(s[i:]+s[:i])
        i += 1
    return permutations


def fibonacci(n):
    a,b = 0,1
    for i in range(n):
        a,b = a+b,a
    return a

# find the nth row in pascal's triangle, useful in combinatorics
def pascals_row(n):
    row = []
    for k in range(n+1):
        elem = math.factorial(n)/(math.factorial(n-k)*math.factorial(k))
        row.append(elem)
    return row

def __prime_factors__(n):
    factors = []
    lastresult = n
    if n <= 0:
        return []
    if n == 1:
        return [1]
    while 1:
        if lastresult == 1:
            break
        c = 2
        while 1:
            if lastresult % c == 0:
                break
            c += 1
        factors.append(c)
        lastresult /= c
    return factors

def __proper_divisors__(n):
    divisors = []
    divisors.append(1)
    for i in xrange(2,int(sqrt(n))+1):
        if not n%i:
            divisors.append(i)
            if not i == n/i:
                divisors.append(n/i)
    return sorted(divisors)

def __is_prime__(n):
    return len(__prime_factors__(n)) == 1

def __is_prime_2__(n):
    if n < 0:
        n = 0 - n
    if n == 0:
        return False
    if n == 1:
        #return True
        return False
    if n == 2:
        return False
    if n == 3:
        return True
    c = 2
    while c < n/2:
        if n % c == 0:
            return False
        c += 1
    return True


#################################################

# Find the sum of all the multiples of 3 or 5 under 1000.
def problem_1():
    j = 0
    for i in xrange(0,1000):
        if not i % 3 or not i % 5:
            j += i
    print j

# Find the sum of the even valued fibonacci numbers under 4,000,000.
def problem_2():
    a = 1
    b = 2
    tmp = sum = 0
    while b < 4000000:
        if not b % 2:
             sum += b
        tmp = a + b
        a = b
        b = tmp
    print sum

# Find the largest prime factor of 600851475143.
def problem_3(num=600851475143):
    for i in xrange(int(math.sqrt(num)), 1, -1):
        if not num % i:
            for j in xrange(int(math.sqrt(i)),1,-1):
                if not i % j:
                    break
            else:
                print i
                break

# Find the largest palindrome number made from the product of two 3-digit numbers.
def problem_4():
    largest_palindrome=0
    for i in xrange(999,0,-1):
        for j in xrange(999,0,-1):
            forward = str(i*j)
            backward = forward[::-1]
            if (forward == backward) and (i*j > largest_palindrome):
                largest_palindrome = i * j
    print largest_palindrome


# Find the smallest positive number that is evenly divisible by all the numbers from 1 to 20.
def problem_5():
    num = 20
    while True:
        for i in xrange(2,21):
            if num % i:
                break
        else:
            print num
            break
        num += 20

# Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
def problem_6():
    sum_of_squares = 0
    square_of_sums = 0
    for i in xrange(1,101):
        sum_of_squares += i**2
    for i in xrange(1,101):
        square_of_sums += i
    square_of_sums *= square_of_sums
    print square_of_sums - sum_of_squares

# Find the 10,001st prime number.
def problem_7():
    "";



# Find the greatest product of five consecutive digits in the 1000-digit number.
def problem_8():
    largest_product = 1
    large_num = "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
    while len(large_num) >= 5:
        product = 1
        first_five = large_num[0:5]
        for c in first_five:
            product *= int(c)
        if product > largest_product:
            largest_product = product
        large_num = large_num[1:]
    print largest_product

# Find the product of the Pythagorean triplet (a,b,c in a**2 + b**2 = c**2) where a + b + c = 1000.
def problem_9():
    "";


# Find the maximum total path from top to bottom of the triangle 
#   same method used on problem 67
#* Note: collapse the triangle from the bottom up, keeping the highest of the two paths
def problem_18():
    fp = open("problem_18.txt")
    triangle = []
    for l in fp.readlines():
        triangle.append(map(int, l.split()))
    fp.close()
    print triangle
    def trunc_last_two(last, secnd):
        sum_row = []
        for i in xrange(len(secnd)):
            sum_row.append(secnd[i] + (last[i] if (last[i] > last[i+1]) else last[i+1]))
        return sum_row
    for i in xrange(len(triangle) - 1):
        triangle.append(trunc_last_two(triangle.pop(), triangle.pop()))
    print triangle



# Find how many Sundays fell on the first of the month during the 20th century (1901/01/01 - 2000/12/31)
def problem_19():
    year = 1901
    month_days = [31,28,31,30,31,30,31,31,30,31,30,31]

    days = 2 # Jan 1, 1901 = Tuesday
    sun_months = 0

    while year < 2001:
        for i in xrange(12):
            days = days % 7
            if not days:
                sun_months += 1
            if i == 1:
                if not year % 4: # need to account for centuries, but not needed here
                    month_days[i] = 29
                else:
                    month_days[i] = 28
            days += month_days[i]
        year += 1
    print sun_months



# Find the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8, and 9.
def problem_24():
    print ''.join(list(itertools.permutations(string.digits))[999999])
    #num = 123456789
    #count = 1
    #while count <= 1000000:
    #    num += 1
    #    if '0' in str(num):
    #        continue
    #    for c in str(num):
    #        if str(num).count(c) > 1:
    #            break
    #    else:
    #        count += 1
    #        if not count % 1000:
    #            print num, count
    #        num += 8
    #print num



# Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
def problem_26(precision=4001):
    decimal.getcontext().prec=precision
    longest = [0, 0]
    for i in xrange(2,1001):
        count = 0
        prevcount = 0
        dec = str(decimal.Decimal(1)/decimal.Decimal(i))[2:-1]
        #print i,dec
        for j in xrange(1, len(dec)):
            if dec.count(dec[j:]) > 1:
                prevcount = count
                count = dec.count(dec[j:])
                if count == prevcount * 2:
                    #print dec[j:], len(dec[j:])
                    if len(dec[j:]) > longest[1]:
                        longest = [i, len(dec[j:])]
    print longest

# 
def problem_27():
    # product = -59231
    max_primes = 0
    product = 0
    max = (0, 0, 0)
    for a in range(-999, 1001):
        for b in range(-999, 1001):
            for n in range(0, 1001):
                if __is_prime_2__(n**2 + a*n + b):
                    if max_primes < n+1:
                        print "(%d**2 + %d*%d + %d) - is prime" % (n, a, n, b)
                        max_primes = n+1
                        product = a * b
                        max = (a, b, n)
                else:
                    break
    print "max primes = {}; product = {}; formula = {}**2 + {}*{} + {}".format(max_primes, product, max[2], max[0], max[2], max[1])

# How many different ways can £2 be made using any number of coins? (1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p))
def problem_31():
    pass

# find the sum of all products whose multiplicand/multiplier/product identity can be written as 1 through 9 pandigital
def problem_32():
    def _pandigital(a, b, c):
        digit = str(a) + str(b) + str(c)
        if len(digit) != 9:
            return False
        for c in digit:
            if c == '0':
                return False
            if digit.count(c) > 1:
                return False
        return True
    sum = 0
    products = []
    for i in xrange(99):
        for j in xrange(9999):
            if _pandigital(i, j, i*j):
                if i*j not in products:
                    print i, j, i*j
                    sum += i*j
                    products.append(i*j)
    print sum
    

# digit cancelling fractions containing two digits in numerator and denominator
def problem_33():
    def _cancel_digits(n, d):
        options = []
        if n >= d:
            return []
        for c in str(n):
            if c in str(d):
                if c == '0':
                    continue
                #print c
                if re.sub(c, '', str(d), 1) != '0':
                    options.append(float(re.sub(c, '', str(n), 1)) / float(re.sub(c, '', str(d), 1)))
        return options
    for i in xrange(11, 99):
        for j in xrange(11, 99):
            #print i, j
            if float(i) / float(j) in _cancel_digits(i,j):
                print "%d / %d" % (i, j)
                

# Find the sum of all numbers which are equal to the sum of the factorial of their digits
def problem_34():
    for i in xrange(9999999):
        sumfact = 0
        for n in str(i):
            sumfact += math.factorial(int(n))
        if sumfact == i:
            print i


# Find all the circular primes below 1000000
# SLOWWWWW @ 5m26.86s
def problem_35():
    primes = sieve_of_eratosthenes_2(1000000)
    for p in primes:
        if '0' in str(p) or '2' in str(p) or '4' in str(p) or '6' in str(p) or '8' in str(p):
            continue
        #if False not in map(__is_prime__, map(int, permurotate(str(p)))):
        #    print p
        for n in permurotate(str(p)):
        #    print n
            if not __is_prime__(int(n)):
                break
        else:
            print p
    return

# Find the sum of the only eleven primes that are both truncatable from left to right and right to left
def problem_37():
    primes = sieve_of_eratosthenes_2(1000000000)
    for p in primes:
        ps = str(p)
        if ps.startswith('1') or ps.endswith('1'):
            continue
        elif '4' in ps or '6' in ps or '8' in ps:
            continue
        elif '2' in ps and not ps.startswith('2'):
            continue
        elif '5' in ps and not ps.startswith('5'):
            continue
        elif len(ps) > 1:
            for l in xrange(1, len(ps)):
                #print ps, l, len(ps)
                #print ps[:l], ps[l:]
                if not __is_prime_2__(int(ps[:l])):
                    break
                if not __is_prime_2__(int(ps[l:])):
                    break
            else:
                print p


# what is the largest 1-9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2,...,n) where n>1?
def problem_38():
    def _pandigital_products(n):
        pdp = '0'
        for i in xrange(1,10):
            s = str(i * n)
            for c in s:
                if c in pdp or s.count(c) > 1 or c == '0':
                    return int(pdp)
            else:
                pdp += s
        return int(pdp)
    largest_pan = 0
    for i in xrange(1,99999):
        pdp = _pandigital_products(i)
        if pdp > largest_pan:
            print i, pdp
            largest_pan = pdp
    print largest_pan

# find the value of p <= 1000 where p is the perimeter of a right angle triangle with the maximum number of solutions of {a,b,c}
def problem_39():
    perimeters = {}
    for i in xrange(1, 501):
        for j in xrange(1, 501):
            sqrt = math.sqrt(i**2 + j**2)
            sum = i+j+sqrt
            if sum < 1001:
                if sum not in perimeters:
                    perimeters[sum] = []
                perimeters[sum].append([i, j, sqrt])
    max_sums = 0
    for p in perimeters:
        if len(perimeters[p]) > max_sums:
            max_sums = len(perimeters[p])
            print p, perimeters[p]

def problem_40():
    dec_string = ""
    i = 0
    while len(dec_string) <= 1000000:
        dec_string += str(i)
        i += 1
    print i
    print int(dec_string[1])*int(dec_string[10])*int(dec_string[100])*int(dec_string[1000])*int(dec_string[10000])*int(dec_string[100000])*int(dec_string[1000000])


# what is the largest n-digit pandigital prime that exists?
def problem_41():
    pass

# coded triangle numbers, how many words, when converted to sums of their numerical positions, are triangle numbers
def problem_42():
    triangles = []
    for i in xrange(1, 21):
        t = (i + 1) * (.5 * i)
        triangles.append(int(t))
    fp = open("problem_42.txt")
    words = fp.read().strip().split(",")
    tws = 0
    for word in words:
        wordsum = 0
        for c in word:
            wordsum += ord(c) - 64
        if wordsum in triangles:
            tws += 1
    print tws

# not fast at all...
def problem_45():
    MAX_I = 100000
    triangles = []
    pentagonals = []
    hexagonals = []
    solutions = []
    for i in xrange(MAX_I):
        triangles.append( i * ( i + 1) / 2 )
        #pentagonals.append( i * ( 3 * i - 1 ) / 2 )
        #hexagonals.append( i * ( 2 * i - 1) )
    print "checking pentagonals"
    for i in xrange(MAX_I):
        p = i * ( 3 * i - 1 ) / 2
        if p in triangles:
            pentagonals.append(p)
    print "checking hexagonals"
    for i in xrange(MAX_I):
        h = i * ( 2 * i - 1)
        if h in pentagonals:
            hexagonals.append(h)
    print hexagonals

#
def problem_46():
    for i in xrange(1, 10000, 2):
        if __is_prime__(i):
            continue
        for j in xrange(int(math.sqrt(i) -1)):
            if __is_prime__(i - (2 * (j**2))):
                print "%d = %d + 2 x %d^2" % (i, i - (2 * (j**2)), j)
                break
        else:
            print i
            break


# How many, not necessarily distinct, values of  nCr, for 1 <= n <= 100, are greater than one-million?
def problem_53():
    count = 0
    for i in range(101):
        for elem in pascals_row(i):
            if elem > 1000000:
                count += 1
    print count


def problem_56():
    #maxAB = ()
    maxVal = 0
    
    for i in xrange(1,101):
        for j in xrange(1,101):
            s = sum(map(int, list(str(i**j))))
            if s > maxVal:
                #maxAB = (i,j)
                maxVal = s
    #print maxAB
    print maxVal


def problem_59():
    cipher = file('problem_59.txt').read().strip().split(',')
    cipher = map(int, cipher)
    #print cipher
    for x in string.lowercase:
        for y in string.lowercase:
            for z in string.lowercase:
                key = [ord(x),ord(y),ord(z)] * 400
                key.append(ord(x))
                decipher = "".join(map(chr, map(lambda x, y: x ^ y, cipher, key)))
                if " the " in decipher:
                    print (x,y,z)
                    print decipher
                    sum = 0
                    for c in decipher:
                        sum += ord(c)
                    print sum

# how many numbers below ten million will have the sum of all digits eventually chain to 89
def problem_92():
    def _sum_of_digits(num):
        sum = 0
        for c in str(num):
            sum += int(c) ** 2
        return sum
    
    sums_89 = 0
    for i in xrange(1,10000001):
        seen_sums = []
        sum = _sum_of_digits(i)
        while sum not in seen_sums:
            if sum == 89:
                sums_89 += 1
                break
            seen_sums.append(sum)
            sum = _sum_of_digits(sum)
    print sums_89


# find the least number for which the proportion of bouncy numbers is 99%
def problem_112():
    def _is_bouncy(num):
        lt = False
        gt = False
        prevc = None
        for c in str(num):
            if prevc == None:
                prevc = int(c)
                continue
            if int(c) > prevc:
                gt = True
            elif int(c) < prevc:
                lt = True
            if lt and gt:
                return True
            prevc = int(c)
        return False
    bouncy_nums = 0
    i = 100
    while 1:
        if _is_bouncy(i):
            bouncy_nums += 1
        if float(i) * 0.99 == float(bouncy_nums):
            break
        i += 1
    print i, bouncy_nums

# how many numbers below a googol (10**100) are not bouncy
def problem_113():
    i = 0
    while i < 10**100:
        i += 1
    pass


#Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0, where each '_' is a single digit
def problem_206():
    for i in xrange(1010101010, 1389026625, 10):
        s = i**2
        if re.match("1.2.3.4.5.6.7.8.9.0", str(s)):
            print "", s, " - ", i

# Find the smallest denominator 'd', having a resilience R(d) < 15499/94744
# arrived at 105380 38240

def problem_243():
    i = 94745
    sum = 0
    while True:
        for j in xrange(1,i):
            if Fraction(j, i)._denominator == i:
                sum += 1
        if Fraction(sum, i-1) < Fraction(15499, 94744):
            print i
            break
        print i, sum
        i += 1
        sum = 0


problem_26()
#sieve_of_eratosthenes_2(1000000)
#print permutate("abcde")
#print map(int, permutate(str(12345)))
#print sieve_of_eratosthenes(1000000)
