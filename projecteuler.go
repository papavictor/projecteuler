/*
  projecteuler.net
  solutions written in Go
  (c)2016
*/

package main

import (
  "fmt"
  "io/ioutil"
  "math"
  "math/big"
  "reflect"
  "sort"
  "strconv"
  "strings"
  "time"
)

// Helper functions
func _build_spiral(side_len int) [][]int {
  large_spiral := make([][]int, side_len)
  if side_len == 1 {
    return [][]int{[]int{1}}
  } else {
    smaller_spiral := _build_spiral(side_len - 1)
    large_spiral = [side_len][side_len]int
    //flip and reverse it, put it at bottom right
    for i:=0; i<len(smaller_spiral); i++ {
      for j:=0; j<len(smaller_spiral); j++ {
        large_spiral[side_len-1-i][side_len-1-j] = smaller_spiral[i][j]
      }
    }
  }
  for i:=0; i<side_len; i++ {
    large_spiral[0][i] = (side_len * (side_len - 1) + 1 + i)
    large_spiral[i][0] = (side_len * (side_len - 1) + 1 - i)
  }
  return large_spiral
}

func _fibonacci(n int) *big.Int {
  a := big.NewInt(0)
  b := big.NewInt(1)
  for i:=0; i<=n; i++ {
    a,b = b, big.NewInt(0).Add(a,b)
  }
  return a
}

func _factorial(n *big.Int) *big.Int {
  if n.Cmp(big.NewInt(0)) == 1 {
    nMinOne := big.NewInt(0).Sub(n, big.NewInt(1))
    childFact := _factorial(nMinOne)
    result := big.NewInt(0).Mul(n, childFact)
    return result
  }
  return big.NewInt(1)
}

func _factors_of(n int) []int {
  factors := []int{}
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n % i == 0 {
      factors = append(factors, i)
      factors = append(factors, _factors_of(n / i)...)
      return factors
    }
  }
  return []int{n}
}

func _is_prime(n int) bool {
  if n == 0 {
    return false
  }
  if n < 0 {
    n = 0 - n
  }
  if n == 1 {
    return false
  } else if n == 2 {
    return true
  } else if n % 2 == 0 {
    return false
  }
  for i:=3; float64(i) <= math.Sqrt(float64(n)); i+=2 {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func _pascals_row(n int) []int {
  row := []int{}
  for i := 0; i <= n; i++ {
    elem := big.NewInt(0)
    a := _factorial(big.NewInt(int64(n)))
    b := _factorial(big.NewInt(int64(n - i)))
    c := _factorial(big.NewInt(int64(i)))
    elem = big.NewInt(0).Div(a, big.NewInt(0).Mul(b, c))
    row = append(row, int(elem.Int64()))
  }
  return row
}

func _permutations(iterable []string, r int, limit int) []string {
  n := len(iterable)
  indices := make([]int, n)
  for i := range indices {
    indices[i] = i
  }
  cycles := make([]int, r)
  for i:= range cycles {
    cycles[i] = n - i
  }
  result := make([]string, r)
  results := make([]string, 0)
  for i, el := range indices[:r] {
    result[i] = iterable[el]
  }
  strresult := strings.Join(result, "")
  results = append(results, strresult)
  if r > n {
    return results
  }
  count := 1
  for n > 0 {
    i := r - 1
    for ; i>=0; i-- {
      cycles[i]--
      if cycles[i] == 0 {
        index := indices[i]
        for j:=i; j<n-1; j++ {
          indices[j] = indices[j+1]
        }
        indices[n-1] = index
        cycles[i] = n - i
      } else {
        count++
        j := cycles[i]
        indices[i], indices[n-j] = indices[n-j], indices[i]
        for k:=i; k<r; k++ {
          result[k] = iterable[indices[k]]
        }
        strresult := strings.Join(result, "")
        results = append(results, strresult)
        break
      }
      if count >= limit {
        break
      }
    }
    if i < 0 || count >= limit {
      break
    }
  }
  return results
}

func _proper_divisors_of(n int) []int {
  divisors := []int{}
  divisors = append(divisors, 1)
  for i:=2; i<=int(math.Sqrt(float64(n))); i++ {
    if n % i == 0 {
      divisors = append(divisors, i)
      if i != n/i {
        divisors = append(divisors, n/i)
      }
    }
  }
  return divisors
}

func _reverse_int(n int) int {
  str := strconv.Itoa(n)
  runestr := []rune(str)
  for i, j := 0, len(runestr) - 1; i <= j; i, j = i+1, j-1 {
    runestr[i], runestr[j] = runestr[j], runestr[i]
  }
  revint, _ := strconv.Atoi(string(runestr))
  return revint
}

func _sieve_of_eratosthenes(n int) []int {
  sieve := []int{2}
  for i := 3; i <= n; i += 2 {
    isPrime := true
    for j := range sieve {
      if sieve[j] > int(math.Sqrt(float64(i))) {
        break
      }
      if i % sieve[j] == 0 {
        isPrime = false
        break
      }
    }
    if isPrime {
      sieve = append(sieve, i)
    }
  }
  return sieve
}

// Problem 1
func problem_1() int {
  sum := 0
  for i := 0; i < 1000; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  return sum
}

// Problem 2
func problem_2() int {
  sum := 0
  for i := 0; _fibonacci(i).Int64() < 4000000; i++ {
    if _fibonacci(i).Int64() % 2 == 0 {
      sum += int(_fibonacci(i).Int64())
    }
  }
  return sum
}

// Problem 3
func problem_3() int {
  factors := _factors_of(600851475143)
  return factors[len(factors) - 1]
}

// Problem 4
func problem_4() int {
  largest_palindrome := 0
  for i := 100; i < 1000; i++ {
    for j := 100; j < 1000; j++ {
      if i * j == _reverse_int(i * j) {
        if i * j > largest_palindrome {
          largest_palindrome = i * j
        }
      }
    }
  }
  return largest_palindrome
}

// Problem 5
func problem_5() int {
  for i := 20; i > 0; i += 20 {
    for j := 2; j <= 20; j++ {
      if i % j != 0 {
        break
      } else if j == 20 {
        return i
      }
    }
  }
  return 0
}

// Problem 6
func problem_6() int {
  sum_of_squares := 0
  square_of_sums := 0
  for i := 1; i <= 100; i++ {
    sum_of_squares += i * i
    square_of_sums += i
  }
  square_of_sums = square_of_sums * square_of_sums
  return square_of_sums - sum_of_squares
}

// Problem 7
func problem_7() int {
  sieve := _sieve_of_eratosthenes(1000000)
  return sieve[10000]
}

// Problem 8
func problem_8() int {
  bignumstr := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
  adjacent_num := 13
  maxproduct := 0
  product := 1
  for i := 0; i < len(bignumstr) - adjacent_num; i++ {
    smallstr := bignumstr[i:i + adjacent_num]
    if strings.Contains(smallstr, "0") {
      continue
    }
    digit := 0
    for j := 0; j < len(smallstr); j++ {
      digit, _ = strconv.Atoi(string(rune(smallstr[j])))
      product = product * digit
    }
    if product > maxproduct {
      maxproduct = product
    }
    product = 1
  }
  return maxproduct
}

// Problem 9
func problem_9() int {
  a, b, c := 1, 2, 3
  for a = 1; a < b; a++ {
    for b = 2; b < c; b++ {
      for c = 3; a + b + c <= 1000; c++ {
        if a + b + c == 1000 {
          if (a * a) + (b * b) == (c * c) {
            return a * b * c
          }
        }
      }
    }
  }
  return 0
}

// Problem 10
func problem_10() int {
  primes := _sieve_of_eratosthenes(2000000)
  sum := 0
  for p := range primes {
    sum += primes[p]
  }
  return sum
}

// Problem 11
func problem_11() int {
  max_product := 0
  max_coords := [2]int{}
  max_divisors := [4]int{}
  num_array := [20][20]int{}
  num_grid := `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`
  num_rows := strings.Split(num_grid, "\n")
  for i := 0; i < len(num_rows); i++ {
    row_cells := strings.Split(num_rows[i], " ")
    for j := 0; j < len(row_cells); j++ {
      digit, _ := strconv.Atoi(row_cells[j])
      num_array[i][j] = digit
    }
  }
  for i := 0; i < 20; i++ {
    for j := 0; j < 20; j++ {
      // test vertical
      product := 0
      if i <= 16 {
        product = num_array[i][j] * num_array[i+1][j] * num_array[i+2][j] * num_array[i+3][j]
        if product > max_product {
          max_product = product
          max_coords[0], max_coords[1] = i, j
          max_divisors[0], max_divisors[1], max_divisors[2], max_divisors[3] = num_array[i][j], num_array[i+1][j], num_array[i+2][j], num_array[i+3][j]
        }
      }
      // test horizontal
      product = 0
      if j <= 16 {
        product = num_array[i][j] * num_array[i][j+1] * num_array[i][j+2] * num_array[i][j+3]
        if product > max_product {
          max_product = product
          max_coords[0], max_coords[1] = i, j
          max_divisors[0], max_divisors[1], max_divisors[2], max_divisors[3] = num_array[i][j], num_array[i][j+1], num_array[i][j+2], num_array[i][j+3]
        }
      }
      // test diagonal (down-right)
      product = 0
      if i <= 16 && j <= 16 {
        product = num_array[i][j] * num_array[i+1][j+1] * num_array[i+2][j+2] * num_array[i+3][j+3]
        if product > max_product {
          max_product = product
          max_coords[0], max_coords[1] = i, j
          max_divisors[0], max_divisors[1], max_divisors[2], max_divisors[3] = num_array[i][j], num_array[i+1][j+1], num_array[i+2][j+2], num_array[i+3][j+3]
        }
      }
      // test diagonal (up-right)
      product = 0
      if i >= 3 && j <= 16 {
        product = num_array[i][j] * num_array[i-1][j+1] * num_array[i-2][j+2] * num_array[i-3][j+3]
        if product > max_product {
          max_product = product
          max_coords[0], max_coords[1] = i, j
          max_divisors[0], max_divisors[1], max_divisors[2], max_divisors[3] = num_array[i][j], num_array[i-1][j+1], num_array[i-2][j+2], num_array[i-3][j+3]
        }
      }
    }
  }
  return max_product
}

// Problem 12
func problem_12() int {
  triangle := 0
  divisors := 2
  for i := 1; divisors < 500; i++ {
    divisors = 2
    triangle = triangle + i
    for j := 2; float64(j) <= math.Sqrt(float64(triangle)); j++ {
      if triangle % j == 0 {
        divisors += 1
        if float64(j) < math.Sqrt(float64(triangle)) {
          divisors += 1
        }
      }
    }
  }
  return triangle
}

// Problem 13
func problem_13() int {
  data, _ := ioutil.ReadFile("problem_13.txt")
  strnumbers := []string{}
  sum := new(big.Int)
  strnumbers = strings.Split(string(data), "\n")
  for i := 0; i < len(strnumbers); i++ {
    num := new(big.Int)
    num.SetString(strnumbers[i], 10)
    sum.Add(sum, num)
  }
  sumstr := sum.String()
  intsum, _ := strconv.Atoi(sumstr[:10])
  return intsum
}

// Problem 14
func problem_14() int {
  longest_chain := 0
  longest_producer := 0
  for i := 3; i < 1000000; i++ {
    chain := 0
    for newi := i; newi != 1; newi = newi {
      if newi % 2 == 0 {
        newi = newi / 2
      } else {
        newi = 3 * newi + 1
      }
      chain += 1
    }
    if chain > longest_chain {
      longest_chain = chain
      longest_producer = i
    }
  }
  return longest_producer
}

// Problem 15
func problem_15() int {
  size_square := 20
  pascals_row := _pascals_row(size_square * 2)
  steps := int(pascals_row[len(pascals_row)/2])
  return steps
}

// Problem 16
func problem_16() int {
  s := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
  sum := 0
  for i:=0; i<len(s); i++ {
    j := int(s[i] - 48)
    sum += j
  }
  return sum
}

// Problem 17
func problem_17() int {
  ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
  teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
  tens := []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
  zeros := []string{"hundred", "thousand"}
  str := ""
  for i:=0; i<=1000; i++ {
    if i < 10 {
      str += ones[i]
    } else if i < 20 {
      str += teens[i%10]
    } else if i < 100 {
      str += tens[i / 10]
      str += ones[i % 10]
    } else if i < 1000 {
      str += ones[i / 100]
      str += zeros[0]
      if i % 100 > 0 {
        str += "and"
      }
      if i % 100 < 20 {
        if i % 100 < 10 {
          str += ones[i % 10]
        } else {
          str += teens[i % 100 % 10]
        }
      } else {
        str += tens[i % 100 / 10]
        str += ones[i % 10]
      }
    } else {
      str += ones[1] + zeros[1]
    }
  }
  return len(str)
}

// Problem 18
func problem_18() int {
  data, _ := ioutil.ReadFile("problem_18.txt")
  pyramid := strings.Split(strings.Trim(string(data), "\n"), "\n")
  lastline := []string{}
  secndlast := []string{}
  origlen := len(pyramid)
  for i:=1; i<origlen; i++ {
    lastline, pyramid = strings.Split(pyramid[len(pyramid)-1], " "), pyramid[:len(pyramid)-1]
    secndlast, pyramid = strings.Split(pyramid[len(pyramid)-1], " "), pyramid[:len(pyramid)-1]
    newline := ""
    for i:= 0; i<len(secndlast); i++ {
      a, _ := strconv.Atoi(secndlast[i])
      b, _ := strconv.Atoi(lastline[i])
      c, _ := strconv.Atoi(lastline[i+1])
      bigV := ""
      if a + b > a + c {
        bigV = strconv.Itoa(a+b)
      } else {
        bigV = strconv.Itoa(a+c)
      }
      newline += bigV + " "
    }
    pyramid = append(pyramid, strings.Trim(newline, " "))
  }
  result, _ := strconv.Atoi(pyramid[0])
  return result
}

// Problem 19
func problem_19() int {
  dpm := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
  dotm := 0
  total_days := 0
  total_sundays := 0
  sunday_adj := 5
  for year:=1901; year<2001; year++ {
    for month:=0; month<12; month++ {
      if month == 1 {
        if year % 400 == 0 {
          dotm = 29
        } else if year % 100 == 0 {
          dotm = dpm[month]
        } else if year % 4 == 0 {
          dotm = 29
        } else {
          dotm = dpm[month]
        }
      } else {
        dotm = dpm[month]
      }
      for i:=0; i<dotm; i++ {
        total_days += 1
        if i == 0 && total_days % 7 == sunday_adj {
          total_sundays += 1
        }
      }
    }
  }
  return total_sundays
}

// Problem 20
func problem_20() int {
  fact := _factorial(big.NewInt(100)).String()
  sum := 0
  for c:=0; c<len(fact); c++ {
    sum += int(fact[c]) - 48
  }
  return sum
}

// Problem 21
func problem_21() int {
  totalsum := 0
  for i:=0; i<10000; i++ {
    divs := _proper_divisors_of(i)
    divsum := 0
    for j:=0; j<len(divs); j++ {
      divsum += divs[j]
    }
    revdivs := _proper_divisors_of(divsum)
    revdivsum := 0
    for k:=0; k<len(revdivs); k++ {
      revdivsum += revdivs[k]
    }
    if revdivsum == i && revdivsum != divsum {
      totalsum += divsum
    }
  }
  return totalsum
}

// Problem 22
func problem_22() int {
  data, _ := ioutil.ReadFile("problem_22.txt")
  names := strings.Split(strings.Trim(string(data), "\n\""), "\",\"")
  sort.Strings(names)
  combined_score := 0
  for i:=0; i<len(names); i++ {
    word_score := 0
    for j:=0; j<len(names[i]); j++ {
      word_score += int(names[i][j]-64)
    }
    word_score = word_score * (i+1)
    combined_score += word_score
  }
  return combined_score
}

// Problem 23
func problem_23() int {
  sum_of_two := [28123]bool{}
  abundant_numbers := []int{}
  sum_of_non_sum := 0
  for i:=1; i<28123; i++ {
    divs := _proper_divisors_of(i)
    sum_of_divs := 0
    for j:=0; j<len(divs); j++ {
      sum_of_divs += divs[j]
    }
    if sum_of_divs > i {
        abundant_numbers = append(abundant_numbers, i)
        for k:=0; k<len(abundant_numbers); k++ {
          if i + abundant_numbers[k] < 28123 {
            sum_of_two[i+abundant_numbers[k]] = true
          }
        }
    }
  }
  for l:=0; l<len(sum_of_two); l++ {
    if sum_of_two[l] == false {
      sum_of_non_sum += l
    }
  }
  return sum_of_non_sum
}

// Problem 24
func problem_24() int {
  numbers := []string{"0","1","2","3","4","5","6","7","8","9"}
  perms := _permutations(numbers, 10, 1000000)
  result, _ := strconv.Atoi(perms[len(perms)-1])
  return result
}

// Problem 25
func problem_25() int {
  i,j := 0,0
  for ; j < 1000; i++ {
    j = len(_fibonacci(i).String())
  }
  return i
}

// Problem 26
func problem_26() int {
  longest := []int{0,0}
  dec := big.NewRat(int64(1),int64(1))
  for _,i := range _sieve_of_eratosthenes(1000) {
    count := 0
    prevcount := 0
    dec = big.NewRat(int64(1), int64(i))
    strdec := dec.FloatString(4001)
    strdec = strdec[2:len(strdec)-1]
    for j:=0; j<len(strdec)-2; j++ {
      if strings.Count(strdec, strdec[j:len(strdec)-2]) > 1 {
        prevcount = count
        count = strings.Count(strdec, strdec[j:len(strdec)-2])
        if count == prevcount * 2 {
          if len(strdec[j:len(strdec)-2]) > longest[1] {
            longest = []int{i, len(strdec[j:len(strdec)-2])}
          }
        }
      }
    }
  }
  return longest[0]
}

// Problem 27
func problem_27() int {
  num_primes := 0
  product := 0
  for a := -999; a <= 999; a++ {
    for b := -1000; b <= 1000; b++ {
      for n := 0; _is_prime( (n*n)+(a*n)+b ) ; n++ {
        if num_primes < n+1 {
            num_primes = n+1
            product = a * b
        }
      }
    }
  }
  return product
}

// Problem 28
func problem_28() int {
  sum := 1
  for i:=3; i<=1001; i++ {
    if i % 2 == 0 {
      sum += i * i
      sum += (i*i)-(i-1)
      sum += (i*i)-(2*(i-1))
      sum += (i*i)-(3*(i-1))
    }
  }
  return sum
}

// Problem 29
func problem_29() int {
  terms := []*big.Int{}
  tmp := big.NewInt(0)
  for i:=2; i<=100; i++ {
    for j:=2; j<=100; j++ {
      found := false
      for k:=0; k<len(terms); k++ {
        //if terms[k].Cmp(big.NewInt(0).Exp(big.NewInt(int64(i)),big.NewInt(int64(j)),nil)) == 0 {
        if terms[k].Cmp(tmp.Exp(big.NewInt(int64(i)),big.NewInt(int64(j)),nil)) == 0 {
          found = true
          break
        }
      }
      if ! found {
        terms = append(terms, big.NewInt(0).Exp(big.NewInt(int64(i)), big.NewInt(int64(j)), nil))
      }
    }
  }
  return len(terms)
}

// Problem 30
func problem_30() int {
  bigsum := 0
  for i:=2; i<1020825; i++ {
    sum := 0
    for _, j := range strconv.Itoa(i) {
      k, _ := strconv.Atoi(string(j))
      sum += int(math.Pow(float64(k), float64(5)))
    }
    if sum == i {
      bigsum += i
    }
  }
  return bigsum
}

// Main function
func main() {
  keys := []string{}
  problems := map[string]interface{}{"problem_01": problem_1, "problem_02": problem_2, "problem_03": problem_3, "problem_04": problem_4,
                                     "problem_05": problem_5, "problem_06": problem_6, "problem_07": problem_7, "problem_08": problem_8,
                                     "problem_09": problem_9, "problem_10": problem_10, "problem_11": problem_11, "problem_12": problem_12,
                                     "problem_13": problem_13, "problem_14": problem_14, "problem_15": problem_15, "problem_16": problem_16,
                                     "problem_17": problem_17, "problem_18": problem_18, "problem_19": problem_19, "problem_20": problem_20,
                                     "problem_21": problem_21, "problem_22": problem_22, "problem_23": problem_23, "problem_24": problem_24,
                                     "problem_25": problem_25, "problem_26": problem_26, "problem_27": problem_27, "problem_28": problem_28,
                                     "problem_29": problem_29, "problem_30": problem_30}
  for k,_ := range problems {
    keys = append(keys, k)
  }
  sort.Strings(keys)
  start := time.Now()
  solution := 0
  end,_ := time.ParseDuration("0s")
  empty_input := make([]reflect.Value, 0)
  fmt.Println(_build_spiral(10))
  /*for _,k := range keys {
    f := reflect.ValueOf(problems[k])
    start = time.Now()
    solution = int(f.Call(empty_input)[0].Int())
    end = time.Since(start)
    fmt.Printf("%s: %d, elapsed: %s\n", k, solution, end)
  }*/
}
