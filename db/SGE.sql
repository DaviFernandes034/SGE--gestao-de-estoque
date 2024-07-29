use SGE;

CREATE TABLE Categorias(

    ID_Categoria int PRIMARY KEY IDENTITY(1,1),
    nome VARCHAR(15) not NULL
);

create table Produtos(

    ID_Produto int PRIMARY key IDENTITY(1,1),
    ID_Categoria int not null,
    nome VARCHAR(50) not null,
    preco DECIMAL not null,
    lote VARCHAR(20) not null,
    validade DATE not NULL

);


alter table Produtos
add CONSTRAINT fk_categoria
FOREIGN key (ID_Categoria)
REFERENCES Categorias(ID_Categoria)


CREATE table status(
    ID_Status int PRIMARY key IDENTITY(1,1),
    nome VARCHAR(10) not null
);


create TABLE Controle(
    ID_Controle int PRIMARY key IDENTITY(1,1),
    ID_Produto int not null,
    ID_Status int not null,
    quantidade int not null,
    data date not null
);


alter TABLE Controle
add CONSTRAINT fk_produto
FOREIGN key (ID_Produto)
REFERENCES Produtos(ID_Produto);

alter TABLE Controle
add CONSTRAINT fk_status
FOREIGN key (ID_Status)
REFERENCES status(ID_Status);