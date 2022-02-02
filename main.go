package main

import (
	"fmt"
	"unsafe"
)

/*
#include <stdio.h>

struct operacion {
	char* operador;
	int operando1;
	int operando2;
	int total;
	char* excepcion;
};

int calcular(struct operacion *valor) {
	int errores = 0;
	switch (valor->operador[0]){
		case '+':
			valor->excepcion = "Sin errores";
			valor->total = valor->operando1 + valor->operando2;
		break;
		case '-':
			valor->excepcion = "Sin errores";
			valor->total = valor->operando1 - valor->operando2;
		break;
		case '*':
			valor->excepcion = "Sin errores";
			valor->total = valor->operando1 * valor->operando2;
		break;
		case '/':
			if(valor->operando2==0){
				errores++;
				valor->excepcion = "Divisor u operador2 no puede ser cero";
			}
			else{
				valor->excepcion = "Sin errores";
				valor->total = valor->operando1 / valor->operando2;
			}
		break;
		default:
			errores++;
			valor->excepcion = "Operador no valida"; 
		break;
	}
	return errores;
}
*/
import "C"

type operacion struct{
	operador *C.char
	operando1 C.int
	operando2 C.int
	total C.int
	excepcion *C.char 
}


func main() {

	datos := &operacion {
		operador: C.CString("+"),
		operando1: C.int(10),
		operando2: C.int(5),
	}

	cerrores := C.calcular( &(*((*C.struct_operacion) (unsafe.Pointer(datos)))) )

	errores := int(cerrores)

	if (errores == 0) {
		fmt.Println(int(datos.operando1),C.GoString(datos.operador),int(datos.operando2),"=",int(datos.total))
	} else {
		fmt.Println(C.GoString(datos.excepcion))
	}

}
