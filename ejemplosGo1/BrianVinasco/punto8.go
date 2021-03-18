package main
import "fmt"

void cargar(char productos[5][41],float precios[5])
{
    int f;
    for(f=0;f<5;f++)
    {
        printf("Ingrese el nombre del producto:");
        gets(productos[f]);
        printf("Ingrese su precio:");
        scanf("%f",&precios[f]);
        fflush(stdin);
    }
}

void precioMayorPrimero(float precios[5])
{
    int f;
    int cant=0;
    for(f=1;f<5;f++)
    {
        if (precios[f]>precios[0])
        {
            cant++;
        }
    }
    printf("La cantidad de productos con un precio mayor al primero ingresado son:%i",cant);
}

void imprimir(char productos[5][41],float precios[5])
{
    int f;
    for(f=0;f<5;f++)
    {
        printf("Producto: %s - Precio: %0.2f\n",productos[f],precios[f]);
    }
}

int main()
{
    char productos[5][41];
    float precios[5];
    cargar(productos,precios);
    imprimir(productos,precios);
    precioMayorPrimero(precios);
    getch();
    return 0;
}