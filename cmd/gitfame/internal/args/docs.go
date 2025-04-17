package args

var flagDocRepository string = "путь до Git репозитория; по умолчанию текущая директория"
var flagDocRevision string = "указатель на коммит; HEAD по умолчанию"
var flagDocOrderBy string = "ключ сортировки результатов; один из lines (дефолт), commits, files.\n" +
	"По умолчанию результаты сортируются по убыванию ключа (lines, commits, files). При равенстве ключей выше будет автор с лексикографически меньшим именем. При использовании флага соответствующее поле в ключе перемещается на первое место."
var flagDocUseCommitter string = "булев флаг, заменяющий в расчётах автора (дефолт) на коммиттера"
var flagDocFormat string = "формат вывода; один из tabular (дефолт), csv, json, json-lines;"
var flagDocExtensions string = "список расширений, сужающий список файлов в расчёте; множество ограничений разделяется запятыми, например, '.go,.md'"
var flagDocLanguages string = "список языков (программирования, разметки и др.), сужающий список файлов в расчёте; множество ограничений разделяется запятыми, например 'go,markdown'"
var flagDocExclude string = "набор Glob паттернов, исключающих файлы из расчёта, например 'foo/*,bar/*'"
var flagDocRestrictTo string = "набор Glob паттернов, исключающий все файлы, не удовлетворяющие ни одному из паттернов набора"
var flagDocProgressBar string = "Показывать прогресс-бар"
var flagDocWorkersCount string = "Количество воркеров для параллельной обработки"
