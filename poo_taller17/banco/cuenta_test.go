package banco

import (
	"testing"
)

// TestDepositarYSaldo prueba que Depositar funcione y ObtenerSaldo devuelva lo correcto
func TestDepositarYSaldo(t *testing.T) {
	// 1. CREAR: Una cuenta nueva
	cuenta := NuevaCuenta("Ana Pérez")

	// 2. VERIFICAR: El saldo inicial es 0
	saldoInicial := cuenta.ObtenerSaldo()
	if saldoInicial != 0 {
		t.Errorf("Saldo inicial debería ser 0, pero es %f", saldoInicial)
	}

	// 3. DEPOSITAR: 100 pesos
	cuenta.Depositar(100)

	// 4. VERIFICAR: El saldo ahora es 100
	saldoFinal := cuenta.ObtenerSaldo()
	if saldoFinal != 100 {
		t.Errorf("Saldo después de depositar 100 debería ser 100, pero es %f", saldoFinal)
	}
}

// TestDepositarMontoNegativo prueba que no se pueda depositar dinero negativo
func TestDepositarMontoNegativo(t *testing.T) {
	cuenta := NuevaCuenta("Carlos Ruiz")

	// Intentar depositar -50 (negativo)
	cuenta.Depositar(-50)

	// El saldo debería seguir siendo 0
	saldo := cuenta.ObtenerSaldo()
	if saldo != 0 {
		t.Errorf("Depositar -50 no debería cambiar el saldo, pero es %f", saldo)
	}
}

// TestMultiplesDepositos prueba varios depósitos seguidos
func TestMultiplesDepositos(t *testing.T) {
	cuenta := NuevaCuenta("María López")

	cuenta.Depositar(100)
	cuenta.Depositar(50)
	cuenta.Depositar(25)

	saldo := cuenta.ObtenerSaldo()
	if saldo != 175 {
		t.Errorf("100 + 50 + 25 debería ser 175, pero es %f", saldo)
	}
}

// Puede usar cosas PRIVADAS porque estamos en el mismo paquete
func TestAccesoPrivado(t *testing.T) {
	t.Run("Puedo usar función privada nuevaCuentaPrivada", func(t *testing.T) {
		// FUNCIONA: nuevaCuentaPrivada es privada pero estamos en el mismo paquete
		cuenta := nuevaCuentaPrivada("Ana Interna")

		// FUNCIONA: saldo es privado pero podemos acceder
		if cuenta.saldo != 100 {
			t.Errorf("Esperaba 100, obtuve %f", cuenta.saldo)
		}
	})

	t.Run("Puedo usar método privado calcularInteres", func(t *testing.T) {
		cuenta := NuevaCuenta("Carlos")
		cuenta.Depositar(1000)

		// FUNCIONA: calcularInteres es privado pero estamos en el mismo paquete
		interes := cuenta.calcularInteres(0.10)
		if interes != 100 {
			t.Errorf("10% de 1000 debería ser 100, obtuve %f", interes)
		}
	})

	t.Run("Puedo modificar campo privado directamente", func(t *testing.T) {
		cuenta := NuevaCuenta("María")

		// FUNCIONA: modifico saldo directamente (es privado pero estoy en el paquete)
		cuenta.saldo = 999

		if cuenta.ObtenerSaldo() != 999 {
			t.Errorf("Esperaba 999, obtuve %f", cuenta.ObtenerSaldo())
		}
	})
}

// Prueba normal con API pública (también funciona)
func TestDepositarPublico(t *testing.T) {
	cuenta := NuevaCuenta("Ana")
	cuenta.Depositar(100)

	if cuenta.ObtenerSaldo() != 100 {
		t.Errorf("Esperaba 100, obtuve %f", cuenta.ObtenerSaldo())
	}
}
