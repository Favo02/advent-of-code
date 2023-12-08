import scala.collection.mutable.ListBuffer
import java.io.{BufferedReader, InputStreamReader}

object Day01 extends App {
  val lines = ListBuffer[String]()

  val reader = new BufferedReader(new InputStreamReader(System.in))
  var line : String = null
  while ({ line = reader.readLine(); line != null }) {
    lines += line
  }
  reader.close()

  def replacePart2(line: String): String = {
    val modifiedLine = line
      .replace("one", "o1e")
      .replace("two", "t2o")
      .replace("three", "t3e")
      .replace("four", "f4r")
      .replace("five", "f5e")
      .replace("six", "s6x")
      .replace("seven", "s7n")
      .replace("eight", "e8t")
      .replace("nine", "n9e")
    modifiedLine
  }

  val part1 = (
    lines.map(_.filter(_.isDigit)(0).toInt - '0'),
    lines.map(_.filter(_.isDigit).takeRight(1).toInt)
  ) .zipped
    .toList
    .map({ case (f, l) => f*10 + l })
    .reduceLeft((el, acc) => acc + el)

  val part2 = (
    lines.map(replacePart2).map(_.filter(_.isDigit)(0).toInt - '0'),
    lines.map(replacePart2).map(_.filter(_.isDigit).takeRight(1).toInt)
  ) .zipped
    .toList
    .map({ case (f, l) => f*10 + l })
    .reduceLeft((el, acc) => acc + el)

  println("Part 1: ", part1)
  println("Part 2: ", part2)
}
