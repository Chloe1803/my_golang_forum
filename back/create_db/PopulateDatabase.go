package forum

import (
	"database/sql"
	"fmt"
	"log"
)

func PopulateDatabase() {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données :", err)
	}

	_, err = db.Exec(`
			INSERT INTO categories (category_name) VALUES
					('Cactus/Succulentes'),
					('Plantes Vertes'),
					('Potager'),
					('Permaculture'),
					('Materiels'),
					('Nuisibles'),
					('Insectes'),
					('Plantes médicinales')
	`)
	if err != nil {
		fmt.Println("cat")
		log.Fatal(err)
	}

	_, err = db.Exec(`
			INSERT INTO users (email, username, password) VALUES
			('jean@jean.com', 'jean', 'jeanjean'),
			('michel@michel.com', 'michel', 'michelmichel'),
			('alice@alice.com', 'alice', 'alicealice'),
			('emma@emma.com', 'emma', 'emmaemma'),
			('pierre@pierre.com', 'pierre', 'pierrepierre'),
			('lucie@lucie.com', 'lucie', 'lucielucie'),
			('olivier@olivier.com', 'olivier', 'olivierolivier'),
			('claire@claire.com', 'claire', 'claireclaire'),
			('antoine@antoine.com', 'antoine', 'antoineantoine'),
			('sophie@sophie.com', 'sophie', 'sophiesophie')
	`)
	if err != nil {
		fmt.Println("Utilisateurs")
		log.Fatal(err)
	}

	_, err = db.Exec(`
			INSERT INTO posts (title, content, user_id) VALUES
					('Problème de cochenilles', 'Bonjour, mes plantes d intérieur sont infestés de cochenilles, comment est-ce que je 
pourrais m en débarasser ?', 1),
					('Conseils pour l entretien des roses', 'Je suis novice en jardinage, quelqu un a des conseils pour l entretien de
s roses ?', 2),
					('Mon potager en permaculture', 'Je partage mon expérience de jardinage en permaculture dans mon potager.', 3),
					('Les escargots mangent tous !', 'Je ne sais plus quoi faire, les escargots mangent tous mes plans au potager...',
4),
					('Multiplier ses plantes grasses', 'Le saviez-vous ? Si vous arracher une feuille d une plante grasse et que vous 
la poser dans la terre, après quelque semaine vous aurez une nouvelle plante !', 5),
					('Recommandation arrosage automatique', 'Je voudrais installer un système d arrosage automatique dans mon jardin, 
auriez-vous des conseils à me donner à ce sujet ?', 6),
					('Plantes aromatiques et médicinales', 'Je voudrais installer dans un coin de mon jardin une sélection de plantes 
médicinales, des recommandations ?', 7),
					('Echanges plantes grasses', 'j ai beaucoup de boutures de plantes grasses, je n ai plus assez de place donc je su
is prête à la donner ou les echanger', 8),
					('Conseils entretiens Monstera', 'Je viens de faire l aquisition d une Monstera mais je ne sais pas comment m en o
ccuper...', 9),
					('Légumes faciles', 'J aimerai faire un petit potager mais je n ai pas la main verte... Est-ce que vous pouvez me 
recommander des legumes faciles à faire pousser ?', 10),
					('C est quoi la permaculture ?', 'J en entend beaucoup parler mais c est quoi exactement la permaculture ?', 1),
					('Les taupes foutent le bordel', 'Les taupes font plein de mottes de terre dans mon jardin, ça m énerve !!', 9),
					("Mes concombres ont un goût amer", "J ai mis un plan de concombres au potager mais ils sont immangeables, ils son
t très amer. Est-ce que vous savez pourquoi ?", 2),
					("Comment faire ses propres semis", "Jaimerai faire mes propres semis, comment faut-il si prendre ?", 6),
					("A quelle fréquences faut-il arroser un cactus ?", "J ai un petit cactus, j ai peur de le tuer, à quelle fréquenc
e dois-je l arroser ?", 9)
	`)
	if err != nil {
		fmt.Println("Posts")
		log.Fatal(err)
	}

	_, err = db.Exec(`
			INSERT INTO comments (post_id, user_id, content) VALUES
			(1,2, 'As-tu essayer de passer tes plantes sous un puissant jet d eau ? Les cochenilles n aiment pas l eau'),
			(1,3, 'Tu peux également essayer un mélange d eau et de savon pour les éliminer.'),
			(2, 1, 'A quelle fréquence tailles-tu tes rosiers ?'),
			(3, 8, 'Est-ce que tu trouves que ton potager est plus productif avec des methodes permaculturelles ?'),
			(3, 3, 'ah oui j ai facilement doublé mes rendements !'),
			(4, 9, 'Malheureusement je ne connais pas de remède miracle. J ai essayé de nombreuses mthodes naturelles mais je n ai rie
n trouvé d aussi efficace que le poisons acheté en jardinerie' ),
			(5, 1, 'Je ne connaissais pas cette technique, je vais essayer !'),
			(5, 4, 'Je vole tous le temps des feuilles de plantes grasses en jardinerie pour agrandir gratuitement ma collection !'),
			(6, 5, 'J ai trouvé un système goute à goute pas chère chez Lidl mais le programmateur est vendu séparement'),
			(7, 6, 'Tu as les classiques lavandes, sauges, citronnelle...'),
			(7, 2, 'Il y aussi la camomille, la mélisse et la menthe'),
			(8, 10, 'ça m interesse !'),
			(9, 3, 'La Monstera, également connue sous le nom de plante à fromage suisse, prospère dans un endroit lumineux indirect, 
nécessite un arrosage régulier mais laissez le sol sécher entre les arrosages, et bénéficie d une humidité ambiante accrue grâce à des pul
vérisations d eau occasionnelles'),
			(9, 1, 'J ai tué la mienne avec un arrosage excessif'),
			(10, 3, 'Si tu as de la place, il y a les courgettes !'),
			(11, 3, 'La permaculture est une approche de conception durable qui s inspire des écosystèmes naturels pour créer des syst
èmes agricoles et de vie humaine harmonieux et durables.'),
			(12, 1, 'As-tu essayer les pièges à taupes ?'),
			(12, 2, 'Voyons Antoine, surveilles ton language !'),
			(12, 9, 'TG Michel'),
			(12, 2, 'Je vais signaler ton post, tu vas te faire ban'),
			(12, 9, 'Ouh j ai peur'),
			(13, 4, 'C est sans doute un manque d eau, les concombres ont besoin d être arroser tous les jours'),
			(14, 7, 'Cela dépend du légumes ou du fruit, il y a différentes méthodes'),
			(15, 8, 'Jarrose les miens une fois tous les quinze jours entre mars et octobre, après je ne les arrose plus pendant l hiv
er')
        `)
	if err != nil {
		fmt.Println("Commentaire")
		log.Fatal(err)
	}

	_, err = db.Exec(`
	UPDATE posts
    SET comments = CASE
        WHEN post_id = 1 THEN 2
        WHEN post_id = 2 THEN 1
		WHEN post_id = 3 THEN 2
		WHEN post_id = 4 THEN 1
		WHEN post_id = 5 THEN 2
		WHEN post_id = 6 THEN 1
		WHEN post_id = 7 THEN 2
		WHEN post_id = 8 THEN 1
		WHEN post_id = 9 THEN 2
		WHEN post_id = 10 THEN 1
		WHEN post_id = 11 THEN 1
		WHEN post_id = 12 THEN 5
		WHEN post_id = 13 THEN 1
		WHEN post_id = 14 THEN 1
		WHEN post_id = 15 THEN 1
        ELSE comments
    END
`)
	if err != nil {
		fmt.Println("Comments")
		log.Fatal(err)
	}

	_, err = db.Exec(`
			INSERT INTO post_category (post_id, category_id) VALUES
			(1,2), (1,6), (1,7),
			(2,2), (2,3), (2,4),
			(3,3), (3,4), (4,3),
			(4,6), (4,7), (5, 1),
			(6, 5), (7, 8), (8, 1),
			(9, 2), (10, 3), (11, 4),
			(12, 6), (13, 3), (14, 3),
			(15, 1)
					`)
	if err != nil {
		fmt.Println("Categorie")
		log.Fatal(err)
	}

}
