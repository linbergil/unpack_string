## «Распаковка строки»

Необходимо написать Go функцию, осуществляющую примитивную распаковку строки,
содержащую повторяющиеся символы/руны, например:
* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "3abc" => "" (некорректная строка)
* "45" => "" (некорректная строка)
* "aaa10b" => "" (некорректная строка)
* "aaa0b" => "aab"
* "" => ""
* "d\n5abc" => "d\n\n\n\n\nabc"

Как видно из примеров, разрешено использование цифр, но не чисел.

В случае, если была передана некорректная строка, функция должна возвращать ошибку.
При необходимости можно выделять дополнительные функции / ошибки.

**(*) Дополнительное задание: поддержка экранирования через `\`:**

**(обратите внимание на косые кавычки)**
* \`qwe\4\5\` => "qwe45"
* \`qwe\45\` => "qwe44444"
* \`qwe\\\5\` => \`qwe\\\\\\\\\\`
* \`qw\ne\`  => "" (некорректная строка)

Как видно из примера, заэкранировать можно только цифру или слэш.