create table quotes(
	"id" serial primary key,
	"text" varchar,
	"author" varchar
);

insert into quotes ("text", "author") values
('Be yourself; everyone else is already taken.', 'Oscar Wilde'),
('Two things are infinite: the universe and human stupidity; and Im not sure about the universe.', 'Albert Einstein'),
('Don’t walk in front of me… I may not follow. Don’t walk behind me… I may not lead. Walk beside me… just be my friend.', 'Albert Camus');