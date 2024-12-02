from sympy import symbols, nonlinsolve

rx, ry, rz, ux, uy, uz, k0, k1, k2 = symbols("rx ry rz ux uy uz k0 k1 k2")

a = nonlinsolve([
  k0 >= 0,
  k1 >= 0,
  k2 >= 0,

  246694783951603 + k0 *  54 - rx - k0 * ux,
  201349632539530 + k0 * -21 - ry - k0 * uy,
  307741668306846 + k0 *  12 - rz - k0 * uz,

  220339749104883 + k1 *  77 - rx - k1 * ux,
  131993821472398 + k1 *   7 - ry - k1 * uy,
  381979584524072 + k1 * -58 - rz - k1 * uz,

  148729713759711 + k2 * 238 - rx - k2 * ux,
  225554040514665 + k2 *  84 - ry - k2 * uy,
   96860758795727 + k2 * 360 - rz - k2 * uz

], (rx, ry, rz, ux, uy, uz, k0, k1, k2))

print(a)

# 270392223533307 + 463714142194110 + 273041846062208

# 246694783951603, 201349632539530, 307741668306846 @ 54, -21, 12
# 220339749104883, 131993821472398, 381979584524072 @ 77, 7, -58
# 148729713759711, 225554040514665, 96860758795727 @ 238, 84, 360

# FiniteSet((26*rz/53 + 7231699849647863/53, 114953700582878678/53 - 331*rz/53, rz, 26*uz/53, -331*uz/53, Complement({uz}, {0}), 846337127918, 981421067224, 573879763083, 866877808876, 876275561042, 579647821655, -rz/uz + 317897713841862/uz, -rz/uz + 325057162625080/uz, -rz/uz + 303457473505607/uz, -rz/uz + 318986369932636/uz, -(rz - 319484450797434)/uz, -(rz - 303763180609923)/uz))