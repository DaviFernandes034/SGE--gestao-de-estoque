create database SGE;

use SGE;

CREATE TABLE Categorias(

    ID_Categoria int PRIMARY KEY IDENTITY(1,1),
    nome VARCHAR(15) not NULL
);
 go

create table Produtos(

    ID_Produto int PRIMARY key IDENTITY(1,1),
    ID_Categoria int not null,
	ID_Status int not null,
    nome VARCHAR(50) not null,
    preco DECIMAL not null,
    lote VARCHAR(20) not null,
    validade DATE not NULL

);


CREATE table status(
    ID_Status int PRIMARY key IDENTITY(1,1),
    nome VARCHAR(10) not null
);


alter table Produtos
add CONSTRAINT fk_categoria
FOREIGN key (ID_Categoria)
REFERENCES Categorias(ID_Categoria);

alter table Produtos
add constraint fk_status
foreign key (ID_Status)
references status(ID_Status);