package users

func CreateTable() string {
	return `
		create table user
		(
			id       int auto_increment primary key,
			user     varchar(200) null,
			email    varchar(200) not null,
			password varchar(200) not null,
		    access   int default 0 not null
		);
`
}
