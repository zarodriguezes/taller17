package banco

// Cuenta representa una cuenta bancaria
type Cuenta struct {
	Nombre string  // Público
	saldo  float64 // Privado (minúscula)
}

// Nuevo crea una nueva cuenta (función pública)
func NuevaCuenta(nombre string) *Cuenta {
	return &Cuenta{
		Nombre: nombre,
		saldo:  0,
	}
}

// Depositar agrega dinero a la cuenta (público)
func (c *Cuenta) Depositar(monto float64) {
	if monto > 0 {
		c.saldo += monto
	}
}

// ObtenerSaldo devuelve el saldo actual (público)
func (c *Cuenta) ObtenerSaldo() float64 {
	return c.saldo
}

func (c *Cuenta) calcularInteres(tasa float64) float64 {
	return c.saldo * tasa
}

// PRIVADO - Solo se usa internamente
func nuevaCuentaPrivada(nombre string) *Cuenta {
	return &Cuenta{
		Nombre: nombre,
		saldo:  100, // Saldo inicial para pruebas internas
	}
}
