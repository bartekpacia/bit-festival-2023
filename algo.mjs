// WARUNKI WEJŚCIOWE
// 1. Liczba żył obciążonych (liczba, int)
//
// 2a. Znany prąd obciążenia (odbiornika) (I – prąd obciążenia)
// LUB
// 2b. moc obciążenia (P – czynna, 𝑐𝑜𝑠𝜑 - współczynnik mocy)
//
//
// 3. Temperatura otaczającego powietrza LUB gruntu (T – temperatura, int)nkn

console.log('test')

function calc(I_obl, Temp) /* returns I_ost */ {
  const I_ost = I_obl / (Temp * 0.85)
  return I_ost  
}

export { calc }
