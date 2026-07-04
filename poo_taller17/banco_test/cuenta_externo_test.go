// banco/cuenta_externo_test.go
package banco_test // ← PAQUETE DIFERENTE

import (
	"CuentaCuenta/banco" // Importamos el paquete
	"testing"
)

// NO puede usar cosas PRIVADAS
func TestAccesoPublico(t *testing.T) {
	t.Run("Puedo usar función pública NuevaCuenta", func(t *testing.T) {
		// ✅ FUNCIONA: NuevaCuenta es pública
		cuenta := banco.NuevaCuenta("Ana Externa")

		// ✅ FUNCIONA: Nombre es público
		if cuenta.Nombre != "Ana Externa" {
			t.Errorf("Esperaba 'Ana Externa', obtuve '%s'", cuenta.Nombre)
		}
	})

	t.Run("NO puedo usar función privada nuevaCuentaPrivada", func(t *testing.T) {
		// NO COMPILA: nuevaCuentaPrivada es privada
		// cuenta := banco.nuevaCuentaPrivada("Ana")
		// Error: cannot refer to unexported name banco.nuevaCuentaPrivada

		t.Log("✅ Esta línea no compilaría si la descomentamos")
	})

	t.Run("NO puedo acceder a campo privado saldo", func(t *testing.T) {
		cuenta := banco.NuevaCuenta("Carlos")

		// NO COMPILA: saldo es privado
		// cuenta.saldo = 500
		// Error: cuenta.saldo undefined (cannot refer to unexported field or method saldo)

		// ✅ PERO puedo usar el método público
		cuenta.Depositar(500)
		if cuenta.ObtenerSaldo() != 500 {
			t.Errorf("Esperaba 500, obtuve %f", cuenta.ObtenerSaldo())
		}
	})

	t.Run("NO puedo usar método privado calcularInteres", func(t *testing.T) {
		cuenta := banco.NuevaCuenta("María")
		cuenta.Depositar(1000)

		// NO COMPILA: calcularInteres es privado
		// interes := cuenta.calcularInteres(0.10)
		// Error: cuenta.calcularInteres undefined (cannot refer to unexported method)

		// ✅ Pero puedo calcular el interés usando la API pública
		saldo := cuenta.ObtenerSaldo()
		interesEsperado := saldo * 0.10
		if interesEsperado != 100 {
			t.Errorf("10%% de 1000 debería ser 100, obtuve %f", interesEsperado)
		}
	})
}

// ✅ Prueba de API pública normal
func TestFlujoCompletoPublico(t *testing.T) {
	// Crear cuenta
	cuenta := banco.NuevaCuenta("Usuario Externo")

	// Depositar
	cuenta.Depositar(250)
	cuenta.Depositar(150)

	// Verificar saldo
	if cuenta.ObtenerSaldo() != 400 {
		t.Errorf("Esperaba 400, obtuve %f", cuenta.ObtenerSaldo())
	}

	// Verificar nombre
	if cuenta.Nombre != "Usuario Externo" {
		t.Errorf("Esperaba 'Usuario Externo', obtuve '%s'", cuenta.Nombre)
	}
}
