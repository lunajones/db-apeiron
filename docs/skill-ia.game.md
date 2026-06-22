PROJECT: APEIRON MMO

CONTEXTO GERAL
- MMORPG focado em PvP skill-based.
- Inspiração principal de combate: New World + Dark Souls.
- Não é tab-target.
- Combate action com hitbox real.
- Esquiva com iFrame (invulnerabilidade temporária).
- Timing, spacing e leitura do inimigo são essenciais.
- Poucas skills, porém impactantes.
- Skill do jogador > gear.
- PvP deve sempre ser considerado nas decisões de arquitetura.
- Criaturas devem parecer players e não NPCs scriptados.

OBJETIVO DO COMBATE
- Mais difícil que The Witcher 3.
- Menos previsível.
- Defesa infinita NÃO deve trivializar combate.
- Criaturas devem ter leitura de distância, pressão, bait, flanco e erro humano aparente.
- IA deve ser difícil mas justa.
- Não quero mobs que só usam 2 golpes repetidos.

ARQUITETURA DO JOGO
- Backend em Golang.
- Arquitetura feita para escalar MMO desde o início.
- Evitar sugestões de protótipo que exijam grande refatoração depois.
- Preferir arquitetura extensível e production-grade.
- Modularidade alta.
- Sistemas desacoplados.
- Data-driven design.

COMUNICAÇÃO
- Comunicação entre serviços via gRPC.
- Não usar REST como arquitetura principal.
- Tipagem forte via protobuf.
- Contratos claros entre serviços.

ARQUITETURA DE ENTIDADES
- Não criar arquivo/código individual por criatura.
- Criaturas devem ser data-driven.
- Uma estrutura genérica de Creature.
- Dados carregados do banco.
- creature_template = DNA da criatura.
- creature_instance = criatura viva no mundo.

Criaturas devem suportar:
- stats
- movement profile
- combat profile
- behavior profile
- needs
- personality
- skills
- memory/context
- emotions (futuro)

CRIATURAS VIVAS
Criaturas devem possuir:
- fome
- sede
- fadiga
- bladder (mijar)
- bowel (defecar)
- stress
- medo
- curiosidade
- territorialidade
- personalidade

Comportamentos emergentes:
- scratch/coçar
- farejar
- descansar
- procurar água
- procurar comida
- recuar
- flanquear
- circular alvo
- hesitar
- trocar lado
- comportamento de matilha

Movimentação deve:
- parecer orgânica
- respeitar turn rate
- aceleração
- desaceleração
- momentum
- curva natural
- steering behavior
- sem viradas instantâneas robóticas

COMBATE
- Hitbox real.
- Dodge com iFrame.
- Sem target lock obrigatório.
- Leitura de distância importa.
- Espaçamento importa.
- Positioning importa.
- PvP first.
- Criaturas precisam funcionar também em PvP mindset.

TERRAIN / WORLD
- Mundo com relevo e altura.
- Backend deve suportar navegação 3D.
- Considerar navmesh / walkable areas.
- Evitar soluções 2D XY simplistas.
- Arquitetura pronta para mundo persistente.

TECNOLOGIAS
Linguagem:
- Golang 1.26.3

Banco:
- PostgreSQL

Comunicação:
- gRPC
- protobuf

LIBS INSTALADAS
- github.com/jackc/pgx/v5@v5.9.2
- github.com/jackc/pgx/v5/pgxpool@v5.9.2
- github.com/joho/godotenv
- github.com/rs/zerolog
- google.golang.org/grpc
- google.golang.org/protobuf

REGRAS DE RESPOSTA
- Sempre sugerir arquitetura escalável.
- Evitar gambiarra/protótipo descartável.
- Evitar respostas genéricas.
- Considerar MMO production-grade.
- Considerar PvP sempre.
- Considerar performance e milhares de entidades.
- Se houver tradeoff, explicar.
- Pensar como MMO AAA, mas viável para dev solo.