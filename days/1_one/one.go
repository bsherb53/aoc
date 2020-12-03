package __one

import (
	"fmt"
)

/*
--- Day 1: Report Repair ---
After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?

Your puzzle answer was 744475.

--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?

Your puzzle answer was 70276940.


*/

func Do() {
	var a, b int
	var c, d, e int
	var p1, p2 bool
	for i := range data {
		for j := range data {
			if data[i]+data[j] == 2020 && !p1 {
				a = data[i]
				b = data[j]
				p1 = true
			}

			if !p2 {
				for k := range data {
					if data[i]+data[j]+data[k] == 2020 {
						c = data[i]
						d = data[j]
						e = data[k]
						p2 = true
						break
					}
				}
			}
			if p1 && p2 {
				goto DONE
			}
		}
	}
DONE:

	fmt.Println(fmt.Sprintf("\t%v * %v  = %v", a, b, a*b))
	fmt.Println(fmt.Sprintf("\t%v * %v * %v = %v", c, d, e, c*d*e))
}

var data = []int{
	1753,
	1858,
	1860,
	1978,
	1758,
	1847,
	2010,
	1679,
	1222,
	1723,
	1592,
	1992,
	1865,
	1635,
	1692,
	1653,
	1485,
	848,
	1301,
	1818,
	1872,
	1883,
	1464,
	2002,
	1736,
	1821,
	1851,
	1299,
	1627,
	1698,
	1713,
	1676,
	1673,
	1448,
	1939,
	1506,
	1896,
	1710,
	1677,
	1894,
	1645,
	1454,
	1972,
	1687,
	265,
	1923,
	1666,
	1761,
	1386,
	2006,
	1463,
	1759,
	1460,
	1722,
	1670,
	1731,
	1732,
	1976,
	1564,
	1380,
	1981,
	1998,
	1912,
	1479,
	1500,
	167,
	1904,
	1689,
	1810,
	1675,
	1811,
	1671,
	1535,
	1624,
	1638,
	1848,
	1646,
	1795,
	1717,
	1803,
	1867,
	1794,
	1774,
	1245,
	1915,
	1601,
	1656,
	1472,
	1700,
	1887,
	1869,
	1876,
	1561,
	1743,
	1900,
	1574,
	1400,
	1950,
	1893,
	1576,
	1903,
	1747,
	1560,
	1445,
	1652,
	633,
	1970,
	1812,
	1807,
	1788,
	1948,
	1588,
	1639,
	1719,
	1680,
	1773,
	1890,
	1347,
	1344,
	1456,
	1691,
	1842,
	1585,
	1953,
	410,
	1791,
	485,
	1412,
	1994,
	1799,
	1955,
	1554,
	1661,
	1708,
	1824,
	1553,
	1993,
	1911,
	1515,
	1545,
	856,
	1685,
	1982,
	1954,
	1480,
	1709,
	1428,
	1829,
	1606,
	1613,
	1941,
	1483,
	1513,
	1664,
	1801,
	1720,
	1984,
	1575,
	1805,
	1833,
	1418,
	1882,
	1746,
	483,
	1674,
	1467,
	1453,
	523,
	1414,
	1800,
	1403,
	1946,
	1868,
	1520,
	1861,
	1580,
	1995,
	1960,
	1625,
	1411,
	1558,
	1817,
	1854,
	1617,
	1478,
	735,
	1593,
	1778,
	1809,
	1584,
	1438,
	1845,
	1712,
	1655,
	1990,
	1578,
	1703,
	1895,
	1765,
	1572,
}
