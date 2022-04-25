CREATE TABLE Transaksi (
	Id int NOT NULL,
	tanggal_order datetime,
	status_pelunasan varchar(255),
	tanggal_pembayaran datetime,
	PRIMARY KEY(Id)
)

CREATE TABLE Detail_Transaksi (
	Id int NOT NULL,
	Id_transaksi NOT NULL ,
	Harga int,
	Jumlah int,
	Subtotal int,
	PRIMARY KEY (Id),
	FOREIGN KEY (Id_transaksi) REFERENCE Transaksi(Id)
)

INSERT INTO `Transaksi` (`id`, `tanggal_order`, `status_pelunasan`, `tanggal_pembayaran`) VALUES (1, '2020-12-01 11:30:00', 'Lunas', '2020-12-01 12:00:00');

INSERT INTO `Detail_Transaksi` (`id`, `Id_transaksi`, `Harga`, `Jumlah`, `Subtotal`) VALUES (1, 1, '10000', 2, '20000');

SELECT Transaksi.Id, Transaksi.tanggal_order, Transaksi.status_pelunasan AS status, Transaksi.tanggal_pembayaran, Detail_Transaksi.Subtotal , Detail_Transaksi.Jumlah
FROM Transaksi JOIN Detail_Transaksi ON Transaksi.Id = Detail_Transaksi.Id_transaksi

