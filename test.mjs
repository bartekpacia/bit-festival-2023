import { calc } from './algo.mjs'

almostEqual(11, 10, 2) // from 8 to 12

function almostEqual(value, expectedValue, tolerance) {
  const min = expectedValue - tolerance
  const max = expectedValue + tolerance
  return value >= min && value <= max
}

function test(index) {
  console.log(`CASE ${index}: ładowarka 18A na saharze w powietrzu 35°C`)

  const I_obl = 18 // A
  const Temp = 0.94 // współczynnik tolerancji dla 35°C dla kabla w powietrzu
  const I_ost = calc(
    I_obl /* maksymalne obciążenie kabla */,
    Temp /*temperatura otoczenia*/
  )

  const expected_I_ost = 22.53 // tyle wyszło z obliczeń ręcznych
  if (!almostEqual(I_ost, expected_I_ost, 0.5)) {
    console.log(`CASE ${index} FAILED: I_ost is ${I_ost}, + I_ost + 'A'`)
  } else {
    console.log(`CASE ${index} PASSED: I_ost is ${I_ost}, + I_ost + 'A'`)
  }
}

test(1)
