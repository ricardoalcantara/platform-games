START TRANSACTION;

-- Insert data into catalogs
INSERT INTO catalogs (id, name, description) VALUES
(1, 'Minecraft', 'Cria, explora, sobrevive, repete. Obtém o Minecraft: Java Edition e Bedrock Edition para PC como pacote promocional. Podes alternar facilmente entre jogos com o iniciador unificado do Minecraft e ter acesso a cross play com qualquer versão atual do Minecraft.'),
(2, 'Terraria', 'Cave, Lute, Explore, Construa: O próprio mundo está ao seu alcance enquanto você luta pela sobrevivência, fortuna e glória. Você se aventurará em extensões cavernosas em busca de tesouros e matérias-primas para criar equipamentos, maquinários e estéticas em constante evolução? Talvez você prefira buscar inimigos cada vez mais poderosos para testar sua habilidade em combate? Ou quem sabe você decida construir sua própria cidade para abrigar a variedade de aliados misteriosos que pode encontrar em suas jornadas?');
-- Add more catalogs as needed ...

-- Insert data into catalog_images
INSERT INTO catalog_images (catalog_id, url) VALUES
(1, 'https://www.minecraft.net/content/dam/games/minecraft/screenshots/Allay.png'),
(1, 'https://www.minecraft.net/content/dam/games/minecraft/screenshots/Mangrove_House.png'),
(1, 'https://www.minecraft.net/content/dam/games/minecraft/screenshots/Mangrove_Biome_2.png'),
(2, 'https://cdn.cloudflare.steamstatic.com/steam/apps/105600/header.jpg?t=1666290860'),
(2, 'https://cdn.cloudflare.steamstatic.com/steam/apps/105600/ss_ae168a00ab08104ba266dc30232654d4b3c919e5.1920x1080.jpg?t=1666290860'),
(2, 'https://cdn.cloudflare.steamstatic.com/steam/apps/105600/ss_8c03886f214d2108cafca13845533eaa3d87d83f.1920x1080.jpg?t=1666290860');
-- Add more image entries as needed ...

-- Insert data into catalog_tags
INSERT INTO catalog_tags (catalog_id, catalog_tag_type_id) VALUES
(1, 1),
(1, 2),
(1, 14),
(1, 16),
(2, 1),
(2, 2),
(2, 14),
(2, 16);
-- Add more tag entries as needed ...

-- Insert data into catalog_variations
INSERT INTO catalog_variations (id, catalog_id, name, price, active, memory, disk_size, max_players) VALUES
(1, 1, 'Bedrock', 29.99, 1, 8192, 5120, 4),
(2, 1, 'Java', 39.99, 1, 16384, 10240, 8),
(3, 2, 'Variation 3', 19.99, 1, 1024, 1024, 0);
-- Add more variation entries as needed ...

-- Insert data into variation_details
INSERT INTO variation_details (catalog_variation_id, name, description) VALUES
(1, 'Versão', 'Oficial'),
(1, 'Versão', 'Java com suporte a mods'),
(2, 'Versão', 'Vanila');
-- Add more detail entries as needed ...

COMMIT;

ROLLBACK;