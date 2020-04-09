package time

func CreateTable() []string {
	var cmds = []string{`
create table time
(
	id int auto_increment,
	user_id int null,
	time_log float(4,2) not null,
	constraint time_pk
		primary key (id),
	constraint User
		foreign key (user_id) references user (id)
			on delete cascade
);`}
	return cmds
}
