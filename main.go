package main

/**
 * Постановка задачи: вычислить гипотенузу прямоугольного треугольника, согласно формуле Hyp = SQRT(a*a + b*b)
 * Исходные данные взять из входного потока STDIN, результат вывести в выходной поток STDOUT.
 * ----------
 * Решение:
 * Используем горутины и канальное взаимодействие для многопоточного вычисления.
 * 1. Требуется вычисление двух квадратов, соответственно создаем отдельную горутину для этого,
 * 1а. А поскольку числа из канала обязаны приходить парами, то делаем блокировку и монопольный захват канала
 * горутиной на чтение пары чисел и отправку пары результатов в исходящий канал, который также должен
 * писаться и читаться парно.
 * 1б. Раз таких каналов несколько, создаем тип "парный канал" и его пакет.
 * 2. Итого требуется 2 горутины, работающие с парными каналами: возведение в квадрат и сложение.
 * 3. Извлечение корня принимает число из канала и отдает его в STDOUT.
 *
 * 4. Канальные ошибки. Для предотвращения "зависания" горутин на каналах, подключаем пакет context и пробрасываем
 * контекст исполнения в горутины.
 * 5. Для блокировки каналов парного чтения/записи подключаем пакет sync и включаем в структуру парного канала мьютекс.
 * 6. Данные приходят из STDIN, но требуются парами, соответственно создаем поставщика данных в парный канал.
 * 7. Вычисления производятся в горутинах согласовано, требуется сервис, поставляющщий весь пакет горутин
 * согласовано в отдельном пакете.
 * 8. Также нужна фабрика, поставляющая заданное количество сервисов горутин в исполнение. Можно в пакете с сервисом
 * т.к. фабрика связана с самим сервисом.
 * Итого требуется создать 2 пакета: с типом "парный канал" и сервисом вычислений.
 * 9. Поскольку тип парного чтения канала и сервис взаимосвязаны, требуется создание модуля.
 * 10. Тип данных, поставляемый поставщиком может быть различным, поэтому формируем парный канал как принимающий
 * некий базовый интерфейс.
 *
 * @author fvn20201125..20201127(NT)
 */
import (
	"fmt"
	calc "github.com/Arhat109/calcHypotenuse/calcService"
	"time"
)

func main() {
	var num int
	for {
		fmt.Printf("\nЗадайте кол-во сервисов:")
		cnt, err := fmt.Scanf("%d\n", &num)
		if err == nil && cnt == 1 && num > 0 {
			fmt.Printf("Ok.\n")
			break
		}

		fmt.Printf("\ncnt=%d, err=%s, num=%d", cnt, err, num)
		fmt.Printf("\nНе, так не пойдет .. надо целое, больше нуля и желательно не шибко много.. повторная попытка")
	}
	err := calc.CalcFabric(num, "float64")
	time.Sleep(20_000_000_000) // это 20сек в НАНОсекундах .. :) многозадачность КООПЕРАТИВНАЯ!

	//	fmt.Printf("\nmain(): А у меня фсё.. ")
	fmt.Printf("\nmain(): А у меня фсё.. err=%s", err)
}
