import { calc } from './algo.mjs'

function almostEqual(value, expectedValue, tolerance) {
  return x + tolerance >= min && x - tolerance <= max
}

function test(index) {
  console.log(`CASE ${index}: ładowarka 18A na saharze w powietrzu 35°C`)

  I_obl = 18 // A
  Temp = 0.94 // współczynnik tolerancji dla 35°C dla kabla w powietrzu
  I_ost = calc(
    I_obl /* maksymalne obciążenie kabla */,
    Temp /*temperatura otoczenia*/
  )

  expected_I_ost = 22.53 // tyle wyszło z obliczeń ręcznych
  if (!between(I_ost, expected_I_ost, 1)) {
    console.log(`CASE ${index} FAILED: I_ost is ${I_ost}, + I_ost + 'A'`)
  }
}

test(1)
