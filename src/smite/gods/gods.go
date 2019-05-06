package gods

import (
	"encoding/json"
	"io/ioutil"
	G "smite/general"
)

// MenuItemDescription - menu item description represented as int
type MenuItemDescription int

// All currently existing MenuItemDescriptions
const (
	NoMenuItemDescriptionM MenuItemDescription = -1
	AreaM                  MenuItemDescription = 0
	Level5M                MenuItemDescription = 1
	MaxCorpsesM            MenuItemDescription = 2
	AuraRadiusM            MenuItemDescription = 3
	Level1M                MenuItemDescription = 4
	Level9M                MenuItemDescription = 5
	DamageReductionM       MenuItemDescription = 6
	HitPointsM             MenuItemDescription = 7
	MoveSpeedM             MenuItemDescription = 8
	AffectsM               MenuItemDescription = 9
	LandingRangeM          MenuItemDescription = 10
	RangeM                 MenuItemDescription = 11
	DamageM                MenuItemDescription = 12
	WitnessRadiusM         MenuItemDescription = 13
	MaxDistanceM           MenuItemDescription = 14
	RadiusCorpseRadiusM    MenuItemDescription = 15
	DurationM              MenuItemDescription = 16
	BonusPowerM            MenuItemDescription = 17
	AttackSpeedM           MenuItemDescription = 18
	RadiusM                MenuItemDescription = 19
	BonusMovementSpeedM    MenuItemDescription = 20
	RangeRadiusM           MenuItemDescription = 21
	Level13M               MenuItemDescription = 22
	ConeM                  MenuItemDescription = 23
	AbilityM               MenuItemDescription = 24
	CorpseLifetimeM        MenuItemDescription = 25
	AirDashRangeM          MenuItemDescription = 26
)

// rawItemPartDescriptionTextToMenuItemDescription - converts menu item text to MenuItemDescription
func rawItemPartDescriptionTextToMenuItemDescription(s string) MenuItemDescription {
	switch s {
	case "Hit Points:":
		return HitPointsM
	case "Level 13:":
		return Level13M
	case "Damage:", "Damage: ", "Damage":
		return DamageM
	case "Max Corpses:":
		return MaxCorpsesM
	case "Corpse Lifetime":
		return CorpseLifetimeM
	case "Witness Radius:":
		return WitnessRadiusM
	case "Area:":
		return AreaM
	case "Radius:":
		return RadiusM
	case "Landing Range:":
		return LandingRangeM
	case "Aura Radius:":
		return AuraRadiusM
	case "Move Speed:":
		return MoveSpeedM
	case "Duration:":
		return DurationM
	case "Bonus Power:":
		return BonusPowerM
	case "Level 9:":
		return Level9M
	case "Ability:":
		return AbilityM
	case "Bonus Movement Speed:":
		return BonusMovementSpeedM
	case "Max Distance:":
		return MaxDistanceM
	case "Attack Speed:":
		return AttackSpeedM
	case "Range:", "Range: ":
		return RangeM
	case "Air Dash Range:":
		return AirDashRangeM
	case "Level 5:":
		return Level5M
	case "Cone:":
		return ConeM
	case "Affects:":
		return AffectsM
	case "Range / Radius:":
		return RangeRadiusM
	case "Level 1:":
		return Level1M
	case "Radius/Corpse Radius:":
		return RadiusCorpseRadiusM
	case "Damage Reduction:":
		return DamageReductionM
	default:
		return NoMenuItemDescriptionM
	}
}

// RankItemDescription - menu item description represented as int
type RankItemDescription int

// All currently existing RankItemDescription
const (
	NoRankItemDescription                  RankItemDescription = -1
	MonolithDuration                       RankItemDescription = 0
	DashMagicalProtectionReductionPerStack RankItemDescription = 1
	ProtectionsperStack                    RankItemDescription = 2
	BonusMoveSpeed                         RankItemDescription = 3
	DurationonWeb                          RankItemDescription = 4
	HighTideSpeed                          RankItemDescription = 5
	HealSteal                              RankItemDescription = 6
	DragonDamagePerTick                    RankItemDescription = 7
	IceSlow                                RankItemDescription = 8
	DebuffLifetime                         RankItemDescription = 9
	MaxPenetration                         RankItemDescription = 10
	LightAllyHeal                          RankItemDescription = 11
	DamageperShot                          RankItemDescription = 12
	ProtectionsShred                       RankItemDescription = 13
	BanishDuration                         RankItemDescription = 14
	ArgusThirdHit                          RankItemDescription = 15
	MovementSpeedperStack                  RankItemDescription = 16
	MinorDamage                            RankItemDescription = 17
	OxDamage                               RankItemDescription = 18
	DamageReflect                          RankItemDescription = 19
	ManaRestored                           RankItemDescription = 20
	InitialFuelCost                        RankItemDescription = 21
	DamageperHit                           RankItemDescription = 22
	DetonationRadius                       RankItemDescription = 23
	FiringDuration                         RankItemDescription = 24
	AdditionalPoisonedDamage               RankItemDescription = 25
	SubmergeDamage                         RankItemDescription = 26
	Goal2                                  RankItemDescription = 27
	Lifetime                               RankItemDescription = 28
	SlowperTick                            RankItemDescription = 29
	ProtectionsReduced                     RankItemDescription = 30
	HealperGod                             RankItemDescription = 31
	SpawnDamage                            RankItemDescription = 32
	WallDuration                           RankItemDescription = 33
	StealthDuration                        RankItemDescription = 34
	EnergyGain                             RankItemDescription = 35
	MinkAttackSpeed                        RankItemDescription = 36
	DamagePerWall                          RankItemDescription = 37
	AttackSpeedDebuff                      RankItemDescription = 38
	PassiveMP5                             RankItemDescription = 39
	MovementSpeedBonus                     RankItemDescription = 40
	ShotDamage                             RankItemDescription = 41
	SubmergeLifetime                       RankItemDescription = 42
	HealperStack                           RankItemDescription = 43
	AmplifiedDamage                        RankItemDescription = 44
	FrostbiteDamageBonus                   RankItemDescription = 45
	Healperpassivestack                    RankItemDescription = 46
	GuardDamage                            RankItemDescription = 47
	PhysicalPowerBuff                      RankItemDescription = 48
	CriticalStrikeChance                   RankItemDescription = 49
	DrunkOMeter                            RankItemDescription = 50
	AxeDamage                              RankItemDescription = 51
	InvulnerabilityDuration                RankItemDescription = 52
	SlashDamage                            RankItemDescription = 53
	BuffLifetime                           RankItemDescription = 54
	DamageperTick                          RankItemDescription = 55
	BerserkDuration                        RankItemDescription = 56
	DetonatedDamage                        RankItemDescription = 57
	DruidAreaDuration                      RankItemDescription = 58
	MovementSpeedBuffDebuff                RankItemDescription = 59
	Manaperdrink                           RankItemDescription = 60
	MinionBonusDamage                      RankItemDescription = 61
	MagicalPowerPerStack                   RankItemDescription = 62
	AttackSpeedConversion                  RankItemDescription = 63
	KappaDamageperHit                      RankItemDescription = 64
	FlagDuration                           RankItemDescription = 65
	SpreadTargetsDuration                  RankItemDescription = 66
	ProtectionsGained                      RankItemDescription = 67
	MeleeBonusDamage                       RankItemDescription = 68
	RootandCrippleDuration                 RankItemDescription = 69
	Rank1                                  RankItemDescription = 70
	Goal3                                  RankItemDescription = 71
	PathDuration                           RankItemDescription = 72
	PhysicalProtectionDebuff               RankItemDescription = 73
	ThreePoisons                           RankItemDescription = 74
	MaxHealsperAbility                     RankItemDescription = 75
	RevivedHealth                          RankItemDescription = 76
	ShieldHealth                           RankItemDescription = 77
	FlatHeal                               RankItemDescription = 78
	SelfHeal                               RankItemDescription = 79
	HealPerHit                             RankItemDescription = 80
	ArgusMovementSpeed                     RankItemDescription = 81
	DamageInitialLanding                   RankItemDescription = 82
	PenetrationperStack                    RankItemDescription = 83
	PhysicalProtectionConversion           RankItemDescription = 84
	DragonProtectionDebuff                 RankItemDescription = 85
	SlowPercent                            RankItemDescription = 86
	EchoBuff                               RankItemDescription = 87
	MovementSpeedperstack                  RankItemDescription = 88
	WinterChill                            RankItemDescription = 89
	Heal                                   RankItemDescription = 90
	MagicalPower                           RankItemDescription = 91
	HeartLifetime                          RankItemDescription = 92
	InitialHeal                            RankItemDescription = 93
	PhysicalDamageMitigation               RankItemDescription = 94
	MaxBuffStacks                          RankItemDescription = 95
	MaxTraps                               RankItemDescription = 96
	GustBonusDamage                        RankItemDescription = 97
	Smashed                                RankItemDescription = 98
	LightCCImmunityDuration                RankItemDescription = 99
	ProtectionsDebuffPerStack              RankItemDescription = 100
	PowerLoss                              RankItemDescription = 101
	MaxRankDebuff                          RankItemDescription = 102
	DecoyLifetime                          RankItemDescription = 103
	MinkDamage                             RankItemDescription = 104
	AdditionalScaling                      RankItemDescription = 105
	SilenceDuration                        RankItemDescription = 106
	NumHitstoFullCharge                    RankItemDescription = 107
	Speed                                  RankItemDescription = 108
	DamageIncrease                         RankItemDescription = 109
	MagicalDamageMitigation                RankItemDescription = 110
	MaxHealthDamage                        RankItemDescription = 111
	LandingDamage                          RankItemDescription = 112
	DamageReduction                        RankItemDescription = 113
	StompDamage                            RankItemDescription = 114
	BuffDuration                           RankItemDescription = 115
	MoveSpeedAxe                           RankItemDescription = 116
	AttackSpeedBuff                        RankItemDescription = 117
	SpearsThrown                           RankItemDescription = 118
	BoostedCC                              RankItemDescription = 119
	TornadoLifetime                        RankItemDescription = 120
	BaseDamageperHit                       RankItemDescription = 121
	MadnessDamagePerHit                    RankItemDescription = 122
	PhysicalProtection                     RankItemDescription = 123
	PullStunDuration                       RankItemDescription = 124
	AttackDamage                           RankItemDescription = 125
	DisarmDuration                         RankItemDescription = 126
	PercentMaxHealthperTick                RankItemDescription = 127
	HealthConversion                       RankItemDescription = 128
	Tier3AcornHeal                         RankItemDescription = 129
	BuffperStack                           RankItemDescription = 130
	StrafePenalty                          RankItemDescription = 131
	BonusEnergyGain                        RankItemDescription = 132
	PhysicalLifestealAxe                   RankItemDescription = 133
	MP5                                    RankItemDescription = 134
	PassiveDuration                        RankItemDescription = 135
	BonusPower                             RankItemDescription = 136
	ConeDamage                             RankItemDescription = 137
	PhysicalLifesteal                      RankItemDescription = 138
	NemesisBuffDuration                    RankItemDescription = 139
	MaxShackles                            RankItemDescription = 140
	PhysicalPenetration                    RankItemDescription = 141
	MaxCharge                              RankItemDescription = 142
	MP5PerStack                            RankItemDescription = 143
	PowerDebuff                            RankItemDescription = 144
	MaxStacks                              RankItemDescription = 145
	TargetDebuffDuration                   RankItemDescription = 146
	Tipsy                                  RankItemDescription = 147
	SpeedSlowDuration                      RankItemDescription = 148
	MesmerizeDuration                      RankItemDescription = 149
	AttackSpeedperStack                    RankItemDescription = 150
	DarkSlow                               RankItemDescription = 151
	HealingDebuff                          RankItemDescription = 152
	FearDurationperHit                     RankItemDescription = 153
	WallPortalRange                        RankItemDescription = 154
	JabDamage                              RankItemDescription = 155
	UppercutDamage                         RankItemDescription = 156
	StunCenter                             RankItemDescription = 157
	CardLifetime                           RankItemDescription = 158
	PoisonSlow                             RankItemDescription = 159
	LightMagicalandPhysicalProtections     RankItemDescription = 160
	BurstDamage                            RankItemDescription = 161
	DistanceScaling                        RankItemDescription = 162
	AbilityBonusDamage                     RankItemDescription = 163
	MP5Increase                            RankItemDescription = 164
	MoveSpeed                              RankItemDescription = 165
	AttackMovementSpeedSlow                RankItemDescription = 166
	AirMovementSpeed                       RankItemDescription = 167
	HighTideDamage                         RankItemDescription = 168
	HealperTick                            RankItemDescription = 169
	SpinDamage                             RankItemDescription = 170
	FissureLifetime                        RankItemDescription = 171
	PhysicalPowerBow                       RankItemDescription = 172
	Damageappliedtwice                     RankItemDescription = 173
	AttackSpeedReduction                   RankItemDescription = 174
	DamageDuration                         RankItemDescription = 175
	HealonKill                             RankItemDescription = 176
	Lifesteal                              RankItemDescription = 177
	ArgusSecondHit                         RankItemDescription = 178
	TrembleDuration                        RankItemDescription = 179
	SpringGrowth                           RankItemDescription = 180
	FumesDuration                          RankItemDescription = 181
	MezDuration                            RankItemDescription = 182
	StealthMovementSpeed                   RankItemDescription = 183
	PowerReduction                         RankItemDescription = 184
	BoarLifetime                           RankItemDescription = 185
	Heat                                   RankItemDescription = 186
	ConeAttackDuration                     RankItemDescription = 187
	DruidHealPerHit                        RankItemDescription = 188
	BasicAttack                            RankItemDescription = 189
	ArgusHeal                              RankItemDescription = 190
	GustsSpawned                           RankItemDescription = 191
	TimeSlowed                             RankItemDescription = 192
	DamageMitigation                       RankItemDescription = 193
	MaxHPDamage                            RankItemDescription = 194
	ProtectionsDebuffLifetime              RankItemDescription = 195
	InnerDamage                            RankItemDescription = 196
	DamagePerHit                           RankItemDescription = 197
	DamageReturnedasHealing                RankItemDescription = 198
	ProtectionsDebuff                      RankItemDescription = 199
	ComboDamage                            RankItemDescription = 200
	TauntDuration                          RankItemDescription = 201
	ColossalDuration                       RankItemDescription = 202
	CooldownperHit                         RankItemDescription = 203
	SpeedBuff                              RankItemDescription = 204
	Damagepertick                          RankItemDescription = 205
	ReturnDamage                           RankItemDescription = 206
	ArgusAuraDamage                        RankItemDescription = 207
	OuterDamage                            RankItemDescription = 208
	MagicalPowerBuff                       RankItemDescription = 209
	CleaveDamage                           RankItemDescription = 210
	CloudDuration                          RankItemDescription = 211
	AirPower                               RankItemDescription = 212
	VisionRange                            RankItemDescription = 213
	ExecuteThreshold                       RankItemDescription = 214
	Drinksperpool                          RankItemDescription = 215
	FirstTargetDuration                    RankItemDescription = 216
	TigerDamage                            RankItemDescription = 217
	MaxTargets                             RankItemDescription = 218
	PhysicalPowerBonus                     RankItemDescription = 219
	ArgusProtections                       RankItemDescription = 220
	DetonationDamage                       RankItemDescription = 221
	HPRestoreMinions                       RankItemDescription = 222
	RootDuration                           RankItemDescription = 223
	MaxAurasPossible                       RankItemDescription = 224
	Slow                                   RankItemDescription = 225
	JealousyDuration                       RankItemDescription = 226
	ArgusSpeed                             RankItemDescription = 227
	ViperShots                             RankItemDescription = 228
	StunDurationAxe                        RankItemDescription = 229
	Healing                                RankItemDescription = 230
	DanceMoveSpeed                         RankItemDescription = 231
	WarriorsWillBoost                      RankItemDescription = 232
	ColossalDurationIncrease               RankItemDescription = 233
	ProtectionsStolen                      RankItemDescription = 234
	DruidHeal                              RankItemDescription = 235
	FirstandSecondCooldownDecreasePerTick  RankItemDescription = 236
	RangedDamage                           RankItemDescription = 237
	TidalMeter                             RankItemDescription = 238
	TwoPoisons                             RankItemDescription = 239
	BoostedProtectionsStolen               RankItemDescription = 240
	ProtectionsBonus                       RankItemDescription = 241
	DamageperDart                          RankItemDescription = 242
	ArgusHealth                            RankItemDescription = 243
	HysteriaPerHit                         RankItemDescription = 244
	Healperdrink                           RankItemDescription = 245
	DamageCenter                           RankItemDescription = 246
	MarkDuration                           RankItemDescription = 247
	PhysicalPowerIncrease                  RankItemDescription = 248
	FullChargeDamage                       RankItemDescription = 249
	SlowDebuff                             RankItemDescription = 250
	PoisonDamage                           RankItemDescription = 251
	Duration                               RankItemDescription = 252
	AttackSpeedIncrease                    RankItemDescription = 253
	DamageHealingperHeart                  RankItemDescription = 254
	DamageperSpear                         RankItemDescription = 255
	PercentofDamageTaken                   RankItemDescription = 256
	Cooldownreduction                      RankItemDescription = 257
	DamageperBlast                         RankItemDescription = 258
	ProtectionDebuff                       RankItemDescription = 259
	WeaveHeal                              RankItemDescription = 260
	DruidSlow                              RankItemDescription = 261
	SummonRange                            RankItemDescription = 262
	PoisonDuration                         RankItemDescription = 263
	AttackProgressionBonus                 RankItemDescription = 264
	StackAuraDuration                      RankItemDescription = 265
	DamageIncreaseperCharge                RankItemDescription = 266
	DamageBuff                             RankItemDescription = 267
	AssaultDamage                          RankItemDescription = 268
	ManaRegenperHeart                      RankItemDescription = 269
	MaxDamage                              RankItemDescription = 270
	Goal1                                  RankItemDescription = 271
	BerserkThreshold                       RankItemDescription = 272
	SickleSlowIncrease                     RankItemDescription = 273
	Protections                            RankItemDescription = 274
	HealingReduction                       RankItemDescription = 275
	ProtectionsStolenPerStack              RankItemDescription = 276
	MadnessDuration                        RankItemDescription = 277
	WebRadius                              RankItemDescription = 278
	AssaultStance                          RankItemDescription = 279
	DamagePerAcorn                         RankItemDescription = 280
	BonusProtections                       RankItemDescription = 281
	DruidDamagePerHit                      RankItemDescription = 282
	DarkMagicalPower                       RankItemDescription = 283
	TauntDurationperHit                    RankItemDescription = 284
	PhysicalPower                          RankItemDescription = 285
	MagicalProtectionDebuff                RankItemDescription = 286
	HealingPerConsumedStack                RankItemDescription = 287
	CooldownReductionGods                  RankItemDescription = 288
	HealingperTick                         RankItemDescription = 289
	AmplifiedSpeed                         RankItemDescription = 290
	WallPortalDuration                     RankItemDescription = 291
	AlliedMitigation                       RankItemDescription = 292
	LightSelfHeal                          RankItemDescription = 293
	PhysicalDamageReduction                RankItemDescription = 294
	Arcane                                 RankItemDescription = 295
	BearStunDuration                       RankItemDescription = 296
	MinionExplode                          RankItemDescription = 297
	MovementSpeed                          RankItemDescription = 298
	DruidPowerDebuff                       RankItemDescription = 299
	Ammo                                   RankItemDescription = 300
	PowerGain                              RankItemDescription = 301
	DamageperCorpse                        RankItemDescription = 302
	WoundDuration                          RankItemDescription = 303
	InitialDamage                          RankItemDescription = 304
	ProtectionsBuff                        RankItemDescription = 305
	PhaseSlow                              RankItemDescription = 306
	BoostedSlow                            RankItemDescription = 307
	ProtectionsDuration                    RankItemDescription = 308
	Power                                  RankItemDescription = 309
	DetonationDuration                     RankItemDescription = 310
	DamageEscalation                       RankItemDescription = 311
	PhysicalPowerperStack                  RankItemDescription = 312
	Stun                                   RankItemDescription = 313
	TigerStun                              RankItemDescription = 314
	Increaseddamagetaken                   RankItemDescription = 315
	SummerHeat                             RankItemDescription = 316
	TurtleDamage                           RankItemDescription = 317
	FistDamage                             RankItemDescription = 318
	FlyingMinionDamage                     RankItemDescription = 319
	BonusPhysicalPower                     RankItemDescription = 320
	ChainSlow                              RankItemDescription = 321
	DashDamage                             RankItemDescription = 322
	LaneMinionDamage                       RankItemDescription = 323
	HealingDebuffDuration                  RankItemDescription = 324
	RangedBonusDamage                      RankItemDescription = 325
	StunDuration                           RankItemDescription = 326
	OtherSourceDamage                      RankItemDescription = 327
	EssenceDrinkerBuff                     RankItemDescription = 328
	BonusDamageperStack                    RankItemDescription = 329
	BlockReflect                           RankItemDescription = 330
	HP5                                    RankItemDescription = 331
	MovementSpeedPerStack                  RankItemDescription = 332
	SelfDamageMitigation                   RankItemDescription = 333
	Silence                                RankItemDescription = 334
	ArgusLifetime                          RankItemDescription = 335
	DropchanceforArrowPickup               RankItemDescription = 336
	Durationonhit                          RankItemDescription = 337
	BasicAttackRangeIncrease               RankItemDescription = 338
	ShatterDamage                          RankItemDescription = 339
	PhysicalProtectionReducedPerStack      RankItemDescription = 340
	MirrorDamage                           RankItemDescription = 341
	Damagewithin55                         RankItemDescription = 342
	CatProtections                         RankItemDescription = 343
	TaoluAssaultBoost                      RankItemDescription = 344
	Shield                                 RankItemDescription = 345
	CarryDuration                          RankItemDescription = 346
	GlyphsSpawned                          RankItemDescription = 347
	PhysicalDamage                         RankItemDescription = 348
	GodDamageperTick                       RankItemDescription = 349
	LightningDamage                        RankItemDescription = 350
	IncreasedHealingperStack               RankItemDescription = 351
	SpeedBuffperTargetShackled             RankItemDescription = 352
	Attackspeedconversion                  RankItemDescription = 353
	ArgusFirstHit                          RankItemDescription = 354
	BleedDuration                          RankItemDescription = 355
	BleedDamage                            RankItemDescription = 356
	HealperMinion                          RankItemDescription = 357
	Range                                  RankItemDescription = 358
	MarkLifetime                           RankItemDescription = 359
	LightHealperTick                       RankItemDescription = 360
	MissingHealthHeal                      RankItemDescription = 361
	DisorientDuration                      RankItemDescription = 362
	AttackRange                            RankItemDescription = 363
	MitigationLost                         RankItemDescription = 364
	CCImmunityDuration                     RankItemDescription = 365
	Bounces                                RankItemDescription = 366
	DamagePerTick                          RankItemDescription = 367
	BenevolenceAura                        RankItemDescription = 368
	DamageAxe                              RankItemDescription = 369
	PullDamageperTick                      RankItemDescription = 370
	DamageTakenIncrease                    RankItemDescription = 371
	StaticDamage                           RankItemDescription = 372
	CCReduction                            RankItemDescription = 373
	AutumnDecay                            RankItemDescription = 374
	MagicalScaling                         RankItemDescription = 375
	GroundSpeed                            RankItemDescription = 376
	BasicAttackDamage                      RankItemDescription = 377
	SlowperStack                           RankItemDescription = 378
	MovementSpeedSlow                      RankItemDescription = 379
	DarkProtectionsDebuff                  RankItemDescription = 380
	AttackSpeedBow                         RankItemDescription = 381
	CritChance                             RankItemDescription = 382
	MaxHalos                               RankItemDescription = 383
	RetrievalCooldownReduction             RankItemDescription = 384
	Charges                                RankItemDescription = 385
	SlowPercentage                         RankItemDescription = 386
	AttackSpeed                            RankItemDescription = 387
	DamageOnHit                            RankItemDescription = 388
	HealperHeart                           RankItemDescription = 389
	ValorAura                              RankItemDescription = 390
	Rank5                                  RankItemDescription = 391
	GustDamage                             RankItemDescription = 392
	TrueDamage                             RankItemDescription = 393
	HPRestoreGods                          RankItemDescription = 394
	CrippleDuration                        RankItemDescription = 395
	Root                                   RankItemDescription = 396
	CavalryChargeBoost                     RankItemDescription = 397
	DarkDamage                             RankItemDescription = 398
	MaximumStacks                          RankItemDescription = 399
	CCRperStack                            RankItemDescription = 400
	MovementSpeedBuff                      RankItemDescription = 401
	HealonAssist                           RankItemDescription = 402
	DamageperWraith                        RankItemDescription = 403
	KillHPThreshold                        RankItemDescription = 404
	BonusSpeedLowHealth                    RankItemDescription = 405
	SoldiersHealth                         RankItemDescription = 406
	Disarm                                 RankItemDescription = 407
	KappaHealth                            RankItemDescription = 408
	LightMovementSpeed                     RankItemDescription = 409
	Broodlings                             RankItemDescription = 410
	BonusDamage                            RankItemDescription = 411
	Hitstodestroy                          RankItemDescription = 412
	AreaDamage                             RankItemDescription = 413
	DetonatedHealing                       RankItemDescription = 414
	DamageBow                              RankItemDescription = 415
	ProtectionBuff                         RankItemDescription = 416
	BounceDamageperTick                    RankItemDescription = 417
	BearSelfBuff                           RankItemDescription = 418
	Damagebeyond55                         RankItemDescription = 419
	ProtectionShred                        RankItemDescription = 420
	ShadowCloneSpeed                       RankItemDescription = 421
	DetonateDamage                         RankItemDescription = 422
	DoTDamage                              RankItemDescription = 423
	IntoxicationDebuffDuration             RankItemDescription = 424
	Radius                                 RankItemDescription = 425
	DamagePerStrike                        RankItemDescription = 426
	Rank3                                  RankItemDescription = 427
	EmergeCrashDamage                      RankItemDescription = 428
	DamagePerBlade                         RankItemDescription = 429
	DruidDamage                            RankItemDescription = 430
	MinionStun                             RankItemDescription = 431
	Mesmerize                              RankItemDescription = 432
	FireDamageperTick                      RankItemDescription = 433
	WeakeningAura                          RankItemDescription = 434
	BonusHealth                            RankItemDescription = 435
	SlamDamageIncrease                     RankItemDescription = 436
	OrbDamage                              RankItemDescription = 437
	HealoverTime                           RankItemDescription = 438
	BasicAttacks                           RankItemDescription = 439
	MovementSpeedIncrease                  RankItemDescription = 440
	BearSlow                               RankItemDescription = 441
	PassivePhysicalPower                   RankItemDescription = 442
	BoostedHealing                         RankItemDescription = 443
	ProtectionReduction                    RankItemDescription = 444
	ProtectionDuration                     RankItemDescription = 445
	KillsforOneStack                       RankItemDescription = 446
	FinalDamage                            RankItemDescription = 447
	KnockupTime                            RankItemDescription = 448
	SlamDamage                             RankItemDescription = 449
	MinionDamage                           RankItemDescription = 450
	TeleportRange                          RankItemDescription = 451
	DashRange                              RankItemDescription = 452
	DebuffDuration                         RankItemDescription = 453
	PolymorphDuration                      RankItemDescription = 454
	MaxChargeDmgHealScale                  RankItemDescription = 455
	Arcs                                   RankItemDescription = 456
	HealthRegeneration                     RankItemDescription = 457
	ShieldDuration                         RankItemDescription = 458
	WebLifetime                            RankItemDescription = 459
	ShadowSlow                             RankItemDescription = 460
	HealingReductionLifetime               RankItemDescription = 461
	LightDamage                            RankItemDescription = 462
	SlowDuration                           RankItemDescription = 463
	SubmergeSlow                           RankItemDescription = 464
	HealPerTick                            RankItemDescription = 465
	AbilityLifesteal                       RankItemDescription = 466
	ConvictionBoost                        RankItemDescription = 467
	LightDuration                          RankItemDescription = 468
	SlowAmount                             RankItemDescription = 469
	Powerperstack                          RankItemDescription = 470
	ThunderDamage                          RankItemDescription = 471
	CatHealth                              RankItemDescription = 472
	CooldownReduction                      RankItemDescription = 473
	HP5Stack                               RankItemDescription = 474
	BonusDamageGods                        RankItemDescription = 475
	SoulHealth                             RankItemDescription = 476
	GodBonusDamage                         RankItemDescription = 477
	DefenseStance                          RankItemDescription = 478
	Damage                                 RankItemDescription = 479
	BearDamage                             RankItemDescription = 480
	Slowduration                           RankItemDescription = 481
	MaximumMitigation                      RankItemDescription = 482
	BearDamagePerSwipe                     RankItemDescription = 483
	BroodlingDamage                        RankItemDescription = 484
	DamagefirstHit                         RankItemDescription = 485
	CrowdControlReduction                  RankItemDescription = 486
	HealthBonus                            RankItemDescription = 487
	GustSlow                               RankItemDescription = 488
	JealousyDamageIncrease                 RankItemDescription = 489
	SubmergeManaRegen                      RankItemDescription = 490
	AdditionalDamage                       RankItemDescription = 491
	DragonBreathDamage                     RankItemDescription = 492
	SecondaryDamage                        RankItemDescription = 493
	DetonateLifetime                       RankItemDescription = 494
	RiftDuration                           RankItemDescription = 495
	NumberConjured                         RankItemDescription = 496
	EmpoweredBuff                          RankItemDescription = 497
	ShackleBonus                           RankItemDescription = 498
	Penetration                            RankItemDescription = 499
	PullDuration                           RankItemDescription = 500
)

// rawItemPartDescriptionTextToMenuItemDescription - converts menu item text to MenuItemDescription
func rawItemPartDescriptionTextToRankItemDescription(s string) RankItemDescription {
	switch s {
	case "Max Halos:":
		return MaxHalos
	case "Movement Speed Increase:":
		return MovementSpeedIncrease
	case "Bonus Health:":
		return BonusHealth
	case "Heal over Time:":
		return HealoverTime
	case "Movement Speed Buff/Debuff:":
		return MovementSpeedBuffDebuff
	case "Slam Damage:":
		return SlamDamage
	case "Amplified Speed:":
		return AmplifiedSpeed
	case "MP5 Increase:":
		return MP5Increase
	case "Broodlings:":
		return Broodlings
	case "Protection Buff:":
		return ProtectionBuff
	case "Protections Shred:":
		return ProtectionsShred
	case "Lightning Damage:":
		return LightningDamage
	case "Physical Protection Reduced Per Stack:":
		return PhysicalProtectionReducedPerStack
	case "Range:":
		return Range
	case "Magical Power:":
		return MagicalPower
	case "Move Speed:":
		return MoveSpeed
	case "Knockup Time:":
		return KnockupTime
	case "Glyphs Spawned:":
		return GlyphsSpawned
	case "Buff Lifetime:":
		return BuffLifetime
	case "Spin Damage:":
		return SpinDamage
	case "Madness Damage Per Hit:":
		return MadnessDamagePerHit
	case "Power Gain:":
		return PowerGain
	case "Card Lifetime:":
		return CardLifetime
	case "Additional Damage:":
		return AdditionalDamage
	case "Druid Slow:":
		return DruidSlow
	case "Shield:":
		return Shield
	case "Damage Reduction:":
		return DamageReduction
	case "Minion Bonus Damage:":
		return MinionBonusDamage
	case "Drinks per pool:":
		return Drinksperpool
	case "Tremble Duration:":
		return TrembleDuration
	case "Detonated Healing:":
		return DetonatedHealing
	case "Light Magical and Physical Protections:":
		return LightMagicalandPhysicalProtections
	case "Penetration per Stack:":
		return PenetrationperStack
	case "Cooldown reduction":
		return Cooldownreduction
	case "1st Target Duration:":
		return FirstTargetDuration
	case "Carry Duration:":
		return CarryDuration
	case "Damage Buff:":
		return DamageBuff
	case "Wound Duration:":
		return WoundDuration
	case "Burst Damage:":
		return BurstDamage
	case "Mink Attack Speed:":
		return MinkAttackSpeed
	case "Power:":
		return Power
	case "Energy Gain:":
		return EnergyGain
	case "Basic Attack Damage:":
		return BasicAttackDamage
	case "Max Rank Debuff:":
		return MaxRankDebuff
	case "Mark Lifetime:":
		return MarkLifetime
	case "Protections:":
		return Protections
	case "Damage Per Tick:":
		return DamagePerTick
	case "Power Debuff:":
		return PowerDebuff
	case "Power Loss:":
		return PowerLoss
	case "Move Speed (Axe):":
		return MoveSpeedAxe
	case "Physical Power Bonus:":
		return PhysicalPowerBonus
	case "Druid Damage Per Hit:":
		return DruidDamagePerHit
	case "Mark Duration:":
		return MarkDuration
	case "Self Heal:":
		return SelfHeal
	case "Heart Lifetime:":
		return HeartLifetime
	case "Dragon Breath Damage:":
		return DragonBreathDamage
	case "Magical Scaling:":
		return MagicalScaling
	case "Bonus Physical Power:":
		return BonusPhysicalPower
	case "Heal per God:":
		return HealperGod
	case "Winter Chill:":
		return WinterChill
	case "Fist Damage:":
		return FistDamage
	case "Mesmerize:":
		return Mesmerize
	case "Jealousy Damage Increase:":
		return JealousyDamageIncrease
	case "Cat Protections:":
		return CatProtections
	case "Shield Duration:":
		return ShieldDuration
	case "Movement Speed per stack:":
		return MovementSpeedperstack
	case "Bonus Speed (Low Health):":
		return BonusSpeedLowHealth
	case "Jealousy Duration:":
		return JealousyDuration
	case "Boosted CC:":
		return BoostedCC
	case "Mesmerize Duration:":
		return MesmerizeDuration
	case "Rank 1:":
		return Rank1
	case "Smashed:":
		return Smashed
	case "Frostbite Damage Bonus:":
		return FrostbiteDamageBonus
	case "Tidal Meter:":
		return TidalMeter
	case "Lifetime:":
		return Lifetime
	case "Magical Power Per Stack:":
		return MagicalPowerPerStack
	case "Heal on Kill:":
		return HealonKill
	case "Pull Duration:":
		return PullDuration
	case "Damage Reflect:":
		return DamageReflect
	case "Detonate Lifetime:":
		return DetonateLifetime
	case "Movement Speed Bonus:":
		return MovementSpeedBonus
	case "Healing:":
		return Healing
	case "Berserk Threshold:":
		return BerserkThreshold
	case "Pull Damage per Tick:":
		return PullDamageperTick
	case "Physical Damage:":
		return PhysicalDamage
	case "Dash Damage:":
		return DashDamage
	case "Healing Debuff:":
		return HealingDebuff
	case "Hits to destroy:":
		return Hitstodestroy
	case "Damage within 55:":
		return Damagewithin55
	case "Physical Power (Bow):":
		return PhysicalPowerBow
	case "Attack Speed per Stack:":
		return AttackSpeedperStack
	case "Initial Heal:":
		return InitialHeal
	case "Emerge/Crash Damage:":
		return EmergeCrashDamage
	case "Outer Damage:":
		return OuterDamage
	case "Tiger Damage:":
		return TigerDamage
	case "Argus Heal:":
		return ArgusHeal
	case "Additional Poisoned Damage:":
		return AdditionalPoisonedDamage
	case "Slow:":
		return Slow
	case "Max Auras Possible:":
		return MaxAurasPossible
	case "Strafe Penalty:":
		return StrafePenalty
	case "Healing Debuff Duration:":
		return HealingDebuffDuration
	case "Soul Health:":
		return SoulHealth
	case "Fumes Duration:":
		return FumesDuration
	case "Berserk Duration:":
		return BerserkDuration
	case "Argus Second Hit:":
		return ArgusSecondHit
	case "Charges:":
		return Charges
	case "Spawn Damage:":
		return SpawnDamage
	case "Buff Duration:":
		return BuffDuration
	case "Uppercut Damage:":
		return UppercutDamage
	case "Penetration:":
		return Penetration
	case "Slow: ":
		return Slow
	case "Fissure Lifetime:":
		return FissureLifetime
	case "Heal:":
		return Heal
	case "Missing Health Heal:":
		return MissingHealthHeal
	case "Speed Buff:":
		return SpeedBuff
	case "Attack Speed Conversion:":
		return AttackSpeedConversion
	case "3 Poisons:":
		return ThreePoisons
	case "Heal Per Tick:":
		return HealPerTick
	case "Argus Aura Damage:":
		return ArgusAuraDamage
	case "Slow duration:":
		return Slowduration
	case "Damage (applied twice):":
		return Damageappliedtwice
	case "Critical Strike Chance:":
		return CriticalStrikeChance
	case "Protections Reduced:":
		return ProtectionsReduced
	case "Autumn Decay:":
		return AutumnDecay
	case "Detonate Damage:":
		return DetonateDamage
	case "Gust Slow:":
		return GustSlow
	case "Attack speed conversion:":
		return Attackspeedconversion
	case "Final Damage:":
		return FinalDamage
	case "Speed Slow Duration:":
		return SpeedSlowDuration
	case "Poison Damage:":
		return PoisonDamage
	case "Bear Self Buff:":
		return BearSelfBuff
	case "MP5:":
		return MP5
	case "Light Damage:":
		return LightDamage
	case "Speed:":
		return Speed
	case "Stealth Duration:":
		return StealthDuration
	case "Max Charge:":
		return MaxCharge
	case "Phase Slow:":
		return PhaseSlow
	case "Drop chance for Arrow Pickup:":
		return DropchanceforArrowPickup
	case "Duration:":
		return Duration
	case "Druid Area Duration:":
		return DruidAreaDuration
	case "Slow Debuff:":
		return SlowDebuff
	case "Argus Lifetime:":
		return ArgusLifetime
	case "Cleave Damage:":
		return CleaveDamage
	case "2 Poisons:":
		return TwoPoisons
	case "Protections Bonus:":
		return ProtectionsBonus
	case "Bleed Damage:":
		return BleedDamage
	case "Tornado Lifetime:":
		return TornadoLifetime
	case "Heal per passive stack:":
		return Healperpassivestack
	case "Rank 3:":
		return Rank3
	case "Cooldown Reduction:":
		return CooldownReduction
	case "Slow Duration":
		return SlowDuration
	case "Boosted Slow:":
		return BoostedSlow
	case "Attack Damage:":
		return AttackDamage
	case "Attack Speed Reduction:":
		return AttackSpeedReduction
	case "Root Duration":
		return RootDuration
	case "Invulnerability Duration:":
		return InvulnerabilityDuration
	case "Attack Speed Buff:":
		return AttackSpeedBuff
	case "Cavalry Charge Boost:":
		return CavalryChargeBoost
	case "Boosted Protections Stolen:":
		return BoostedProtectionsStolen
	case "Num Hits to Full Charge:":
		return NumHitstoFullCharge
	case "Damage Per Blade:":
		return DamagePerBlade
	case "Initial Damage:":
		return InitialDamage
	case "Power Reduction:":
		return PowerReduction
	case "Echo Buff:":
		return EchoBuff
	case "Shot Damage:":
		return ShotDamage
	case "Shield Health:":
		return ShieldHealth
	case "Kappa Damage per Hit:":
		return KappaDamageperHit
	case "Return Damage:":
		return ReturnDamage
	case "Kill HP Threshold:":
		return KillHPThreshold
	case "Colossal Duration:":
		return ColossalDuration
	case "Arcs:":
		return Arcs
	case "Sickle Slow Increase:":
		return SickleSlowIncrease
	case "Max Shackles:":
		return MaxShackles
	case "MP5 Per Stack":
		return MP5PerStack
	case "Slam Damage Increase:":
		return SlamDamageIncrease
	case "Increased damage taken:":
		return Increaseddamagetaken
	case "Max HP Damage:":
		return MaxHPDamage
	case "Dark Slow:":
		return DarkSlow
	case "Increased Healing per Stack:":
		return IncreasedHealingperStack
	case "Minor Damage:":
		return MinorDamage
	case "Movement Speed Buff:":
		return MovementSpeedBuff
	case "Mink Damage:":
		return MinkDamage
	case "Dark Protections Debuff:":
		return DarkProtectionsDebuff
	case "Taunt Duration:":
		return TauntDuration
	case "Silence:":
		return Silence
	case "Light Duration:":
		return LightDuration
	case "Bear Stun Duration:":
		return BearStunDuration
	case "Hysteria Per Hit:":
		return HysteriaPerHit
	case "Flag Duration:":
		return FlagDuration
	case "Magical Damage Mitigation:":
		return MagicalDamageMitigation
	case "Buff per Stack:":
		return BuffperStack
	case "Attack Speed:":
		return AttackSpeed
	case "Summer Heat:":
		return SummerHeat
	case "Damage Increase per Charge:":
		return DamageIncreaseperCharge
	case "Protections per Stack:":
		return ProtectionsperStack
	case "Kappa Health:":
		return KappaHealth
	case "Dash Magical Protection Reduction Per Stack:":
		return DashMagicalProtectionReductionPerStack
	case "Heat:":
		return Heat
	case "Health Bonus:":
		return HealthBonus
	case "Shadow Clone Speed:":
		return ShadowCloneSpeed
	case "Passive Duration:":
		return PassiveDuration
	case "Damage per tick:":
		return Damagepertick
	case "Slow Percentage:":
		return SlowPercentage
	case "Minion Stun:":
		return MinionStun
	case "Fear Duration per Hit:":
		return FearDurationperHit
	case "HP Restore (Gods):":
		return HPRestoreGods
	case "Boar Lifetime:":
		return BoarLifetime
	case "Chain Slow:":
		return ChainSlow
	case "Damage Returned as Healing:":
		return DamageReturnedasHealing
	case "Valor Aura:":
		return ValorAura
	case "Dark Damage:":
		return DarkDamage
	case "Empowered Buff:":
		return EmpoweredBuff
	case "Inner Damage:":
		return InnerDamage
	case "Monolith Duration:":
		return MonolithDuration
	case "Essence Drinker Buff:":
		return EssenceDrinkerBuff
	case "Physical Power Buff:":
		return PhysicalPowerBuff
	case "Goal #3:":
		return Goal3
	case "Minion Explode:":
		return MinionExplode
	case "Bonus Move Speed:":
		return BonusMoveSpeed
	case "Silence Duration:":
		return SilenceDuration
	case "Physical Protection:":
		return PhysicalProtection
	case "Magical Power Buff:":
		return MagicalPowerBuff
	case "Bonus Protections:":
		return BonusProtections
	case "Stealth Movement Speed:":
		return StealthMovementSpeed
	case "Max Stacks:":
		return MaxStacks
	case "Healing per Tick:":
		return HealingperTick
	case "Crit Chance:":
		return CritChance
	case "Thunder Damage:":
		return ThunderDamage
	case "Slow %:":
		return SlowPercent
	case "Axe Damage:":
		return AxeDamage
	case "Heal per Minion:":
		return HealperMinion
	case "Boosted Healing:":
		return BoostedHealing
	case "Gust Damage:":
		return GustDamage
	case "Max Charge Dmg/Heal Scale:":
		return MaxChargeDmgHealScale
	case "Initial Fuel Cost:":
		return InitialFuelCost
	case "Damage Per Strike:":
		return DamagePerStrike
	case "Heal Per Hit:":
		return HealPerHit
	case "DoT Damage:":
		return DoTDamage
	case "Movement Speed:":
		return MovementSpeed
	case "Cripple Duration:":
		return CrippleDuration
	case "Goal #2:":
		return Goal2
	case "Attack Speed Debuff:":
		return AttackSpeedDebuff
	case "Execute Threshold:":
		return ExecuteThreshold
	case "Spring Growth:":
		return SpringGrowth
	case "Damage per Shot:":
		return DamageperShot
	case "Dragon Damage Per Tick:":
		return DragonDamagePerTick
	case "Physical Power per Stack":
		return PhysicalPowerperStack
	case "Damage per Tick:":
		return DamageperTick
	case "Dragon Protection Debuff:":
		return DragonProtectionDebuff
	case "Physical Power Increase:":
		return PhysicalPowerIncrease
	case "Wall Portal Range:":
		return WallPortalRange
	case "Submerge Mana Regen:":
		return SubmergeManaRegen
	case "Tiger Stun:":
		return TigerStun
	case "Rift Duration:":
		return RiftDuration
	case "Duration":
		return Duration
	case "Stun Duration (Axe):":
		return StunDurationAxe
	case "Max Penetration:":
		return MaxPenetration
	case "Ranged Damage:":
		return RangedDamage
	case "Heal per Heart:":
		return HealperHeart
	case "Wall Portal Duration:":
		return WallPortalDuration
	case "Nemesis Buff Duration:":
		return NemesisBuffDuration
	case "Max Damage:":
		return MaxDamage
	case "True Damage:":
		return TrueDamage
	case "Debuff Lifetime:":
		return DebuffLifetime
	case "Power per stack":
		return Powerperstack
	case "Submerge Slow:":
		return SubmergeSlow
	case "Damage Taken Increase:":
		return DamageTakenIncrease
	case "Dark Magical Power:":
		return DarkMagicalPower
	case "Heal per Stack:":
		return HealperStack
	case "Protection Shred:":
		return ProtectionShred
	case "High Tide Speed:":
		return HighTideSpeed
	case "Bonus Damage:":
		return BonusDamage
	case "Max Targets:":
		return MaxTargets
	case "Gusts Spawned:":
		return GustsSpawned
	case "Guard Damage:":
		return GuardDamage
	case "Cloud Duration:":
		return CloudDuration
	case "Stun (Center):":
		return StunCenter
	case "Duration on Web:":
		return DurationonWeb
	case "Druid Power Debuff:":
		return DruidPowerDebuff
	case "Damage (Initial/Landing):":
		return DamageInitialLanding
	case "Damage (Bow):":
		return DamageBow
	case "Argus Health:":
		return ArgusHealth
	case "Weave Heal:":
		return WeaveHeal
	case "Protections Debuff:":
		return ProtectionsDebuff
	case "Argus Movement Speed:":
		return ArgusMovementSpeed
	case "Physical Protection Conversion:":
		return PhysicalProtectionConversion
	case "Allied Mitigation:":
		return AlliedMitigation
	case "Orb Damage:":
		return OrbDamage
	case "Cone Damage:":
		return ConeDamage
	case "Bonus Energy Gain:":
		return BonusEnergyGain
	case "Protections Gained:":
		return ProtectionsGained
	case "CC Immunity Duration:":
		return CCImmunityDuration
	case "Mana Restored:":
		return ManaRestored
	case "Flying Minion Damage:":
		return FlyingMinionDamage
	case "Goal #1:":
		return Goal1
	case "Lifesteal:":
		return Lifesteal
	case "God Bonus Damage:":
		return GodBonusDamage
	case "Healing Reduction":
		return HealingReduction
	case "Max Stacks":
		return MaxStacks
	case "Speed Buff per Target Shackled:":
		return SpeedBuffperTargetShackled
	case "% of Damage Taken:":
		return PercentofDamageTaken
	case "Movement Speed Slow:":
		return MovementSpeedSlow
	case "Assault Stance:":
		return AssaultStance
	case "Poison Slow:":
		return PoisonSlow
	case "Protection Duration:":
		return ProtectionDuration
	case "Conviction Boost:":
		return ConvictionBoost
	case "Protections Debuff Lifetime:":
		return ProtectionsDebuffLifetime
	case "Shatter Damage:":
		return ShatterDamage
	case "Bonus Damage (Gods):":
		return BonusDamageGods
	case "Additional Scaling:":
		return AdditionalScaling
	case "Attack Range:":
		return AttackRange
	case "Web Radius:":
		return WebRadius
	case "Turtle Damage:":
		return TurtleDamage
	case "Slash Damage:":
		return SlashDamage
	case "Maximum Mitigation:":
		return MaximumMitigation
	case "Slow per Stack:":
		return SlowperStack
	case "Protections Duration:":
		return ProtectionsDuration
	case "Disarm Duration:":
		return DisarmDuration
	case "Heal":
		return Heal
	case "Banish Duration:":
		return BanishDuration
	case "Damage first Hit:":
		return DamagefirstHit
	case "Submerge Lifetime:":
		return SubmergeLifetime
	case "Bonus Damage per Stack:":
		return BonusDamageperStack
	case "Ability Lifesteal:":
		return AbilityLifesteal
	case "Damage (Center):":
		return DamageCenter
	case "Basic Attacks:":
		return BasicAttacks
	case "Bounce Damage per Tick:":
		return BounceDamageperTick
	case "Slow Duration:":
		return SlowDuration
	case "Mana Regen per Heart:":
		return ManaRegenperHeart
	case "Max Traps:":
		return MaxTraps
	case "Damage per Blast:":
		return DamageperBlast
	case "Disarm:":
		return Disarm
	case "Crowd Control Reduction:":
		return CrowdControlReduction
	case "Intoxication Debuff Duration:":
		return IntoxicationDebuffDuration
	case "CCR per Stack:":
		return CCRperStack
	case "Target Debuff Duration:":
		return TargetDebuffDuration
	case "Ability Bonus Damage:":
		return AbilityBonusDamage
	case "Radius:":
		return Radius
	case "Rank 5:":
		return Rank5
	case "Attack / Movement Speed Slow:":
		return AttackMovementSpeedSlow
	case "Damage per Spear:":
		return DamageperSpear
	case "Ground Speed:":
		return GroundSpeed
	case "Damage Per Tick":
		return DamagePerTick
	case "Damage (Axe):":
		return DamageAxe
	case "Debuff Duration:":
		return DebuffDuration
	case "Taolu Assault Boost:":
		return TaoluAssaultBoost
	case "Attack Speed Increase:":
		return AttackSpeedIncrease
	case "Number Conjured:":
		return NumberConjured
	case "Damage/Healing per Heart:":
		return DamageHealingperHeart
	case "Damage Duration:":
		return DamageDuration
	case "Viper Shots:":
		return ViperShots
	case "High Tide Damage:":
		return HighTideDamage
	case "Protections Stolen Per Stack:":
		return ProtectionsStolenPerStack
	case "1 and 2 Cooldown Decrease Per Tick:":
		return FirstandSecondCooldownDecreasePerTick
	case "Submerge Damage:":
		return SubmergeDamage
	case "Bounces:":
		return Bounces
	case "Fire Damage per Tick:":
		return FireDamageperTick
	case "Amplified Damage:":
		return AmplifiedDamage
	case "Lane Minion Damage:":
		return LaneMinionDamage
	case "Wall Duration:":
		return WallDuration
	case "Cooldown per Hit:":
		return CooldownperHit
	case "Arcane:":
		return Arcane
	case "HP5 / Stack:":
		return HP5Stack
	case "Shackle Bonus:":
		return ShackleBonus
	case "Damage Per Hit:":
		return DamagePerHit
	case "Damage Mitigation:":
		return DamageMitigation
	case "Physical Lifesteal:":
		return PhysicalLifesteal
	case "Light Heal per Tick:":
		return LightHealperTick
	case "Druid Damage:":
		return DruidDamage
	case "Damage Per Acorn:":
		return DamagePerAcorn
	case "Soldier's Health:":
		return SoldiersHealth
	case "Defense Stance:":
		return DefenseStance
	case "Spears Thrown:":
		return SpearsThrown
	case "Web Lifetime:":
		return WebLifetime
	case "Cone Attack Duration:":
		return ConeAttackDuration
	case "Flat Heal:":
		return FlatHeal
	case "Healing Per Consumed Stack":
		return HealingPerConsumedStack
	case "Protection Debuff:":
		return ProtectionDebuff
	case "Summon Range:":
		return SummonRange
	case "Madness Duration:":
		return MadnessDuration
	case "Bear Damage:":
		return BearDamage
	case "Tipsy:":
		return Tipsy
	case "Detonation Damage:":
		return DetonationDamage
	case "Ammo:":
		return Ammo
	case "Argus Protections:":
		return ArgusProtections
	case "Duration on hit:":
		return Durationonhit
	case "Stack/Aura Duration:":
		return StackAuraDuration
	case "Base Damage per Hit:":
		return BaseDamageperHit
	case "Passive MP5:":
		return PassiveMP5
	case "Damage Per Wall:":
		return DamagePerWall
	case "Assault Damage:":
		return AssaultDamage
	case "Static Damage:":
		return StaticDamage
	case "% Max Health per Tick:":
		return PercentMaxHealthperTick
	case "Detonation Duration:":
		return DetonationDuration
	case "Ranged Bonus Damage":
		return RangedBonusDamage
	case "Heal on Assist:":
		return HealonAssist
	case "Mitigation Lost:":
		return MitigationLost
	case "Vision Range:":
		return VisionRange
	case "Cooldown Reduction (Gods):":
		return CooldownReductionGods
	case "Root:":
		return Root
	case "Melee Bonus Damage:":
		return MeleeBonusDamage
	case "Damage Escalation:":
		return DamageEscalation
	case "Decoy Lifetime:":
		return DecoyLifetime
	case "Physical Damage Mitigation:":
		return PhysicalDamageMitigation
	case "Damage:":
		return Damage
	case "Max Health Damage:":
		return MaxHealthDamage
	case "Warrior's Will Boost:":
		return WarriorsWillBoost
	case "HP Restore (Minions):":
		return HPRestoreMinions
	case "Poison Duration:":
		return PoisonDuration
	case "God Damage per Tick:":
		return GodDamageperTick
	case "Damage On Hit:":
		return DamageOnHit
	case "Heal Steal:":
		return HealSteal
	case "Root Duration:":
		return RootDuration
	case "Stomp Damage:":
		return StompDamage
	case "Air Movement Speed:":
		return AirMovementSpeed
	case "Tier 3 Acorn Heal:":
		return Tier3AcornHeal
	case "Other Source Damage:":
		return OtherSourceDamage
	case "Secondary Damage:":
		return SecondaryDamage
	case "Movement Speed per Stack:":
		return MovementSpeedperStack
	case "Block Reflect:":
		return BlockReflect
	case "Max Buff Stacks:":
		return MaxBuffStacks
	case "Pull Stun Duration:":
		return PullStunDuration
	case "Physical Power:":
		return PhysicalPower
	case "Druid Heal:":
		return DruidHeal
	case "Cat Health:":
		return CatHealth
	case "Bonus Power:":
		return BonusPower
	case "Mez Duration:":
		return MezDuration
	case "Kills for One Stack:":
		return KillsforOneStack
	case "Healing Reduction Lifetime:":
		return HealingReductionLifetime
	case "Stun Duration:":
		return StunDuration
	case "Heal per Tick:":
		return HealperTick
	case "Broodling Damage:":
		return BroodlingDamage
	case "Polymorph Duration:":
		return PolymorphDuration
	case "Maximum Stacks:":
		return MaximumStacks
	case "Mana per drink:":
		return Manaperdrink
	case "Protection Reduction:":
		return ProtectionReduction
	case "Dance Move Speed:":
		return DanceMoveSpeed
	case "Combo Damage:":
		return ComboDamage
	case "Health Regeneration:":
		return HealthRegeneration
	case "Basic Attack Range Increase:":
		return BasicAttackRangeIncrease
	case "Retrieval Cooldown Reduction:":
		return RetrievalCooldownReduction
	case "Damage per Hit:":
		return DamageperHit
	case "Heal per drink:":
		return Healperdrink
	case "Area Damage:":
		return AreaDamage
	case "Argus Third Hit:":
		return ArgusThirdHit
	case "Benevolence Aura:":
		return BenevolenceAura
	case "Self Damage Mitigation:":
		return SelfDamageMitigation
	case "Weakening Aura:":
		return WeakeningAura
	case "Stun:":
		return Stun
	case "Minion Damage:":
		return MinionDamage
	case "Gust Bonus Damage:":
		return GustBonusDamage
	case "Bleed Duration:":
		return BleedDuration
	case "Dash Range:":
		return DashRange
	case "Teleport Range:":
		return TeleportRange
	case "Protections Debuff Per Stack":
		return ProtectionsDebuffPerStack
	case "Passive Physical Power:":
		return PassivePhysicalPower
	case "Detonated Damage:":
		return DetonatedDamage
	case "Argus First Hit:":
		return ArgusFirstHit
	case "Revived Health:":
		return RevivedHealth
	case "Taunt Duration per Hit:":
		return TauntDurationperHit
	case "Mirror Damage:":
		return MirrorDamage
	case "Attack Progression Bonus:":
		return AttackProgressionBonus
	case "Protections Buff:":
		return ProtectionsBuff
	case "HP5:":
		return HP5
	case "Disorient Duration:":
		return DisorientDuration
	case "Root and Cripple Duration:":
		return RootandCrippleDuration
	case "Jab Damage:":
		return JabDamage
	case "Physical Protection Debuff:":
		return PhysicalProtectionDebuff
	case "Bear Damage Per Swipe:":
		return BearDamagePerSwipe
	case "Drunk-O-Meter:":
		return DrunkOMeter
	case "Light Movement Speed:":
		return LightMovementSpeed
	case "Spread Targets Duration:":
		return SpreadTargetsDuration
	case "Argus Speed:":
		return ArgusSpeed
	case "Physical Lifesteal (Axe):":
		return PhysicalLifestealAxe
	case "Full Charge Damage:":
		return FullChargeDamage
	case "Landing Damage:":
		return LandingDamage
	case "Distance Scaling:":
		return DistanceScaling
	case "Ice Slow:":
		return IceSlow
	case "Basic Attack:":
		return BasicAttack
	case "Damage per Corpse:":
		return DamageperCorpse
	case "Druid Heal Per Hit:":
		return DruidHealPerHit
	case "Detonation Radius:":
		return DetonationRadius
	case "Firing Duration:":
		return FiringDuration
	case "Protections Stolen:":
		return ProtectionsStolen
	case "Air Power:":
		return AirPower
	case "CC Reduction:":
		return CCReduction
	case "Bear Slow:":
		return BearSlow
	case "Slow per Tick:":
		return SlowperTick
	case "Healing Reduction:":
		return HealingReduction
	case "Light Ally Heal:":
		return LightAllyHeal
	case "Time Slowed:":
		return TimeSlowed
	case "Movement Speed Per Stack":
		return MovementSpeedPerStack
	case "Health Conversion:":
		return HealthConversion
	case "Physical Damage Reduction:":
		return PhysicalDamageReduction
	case "Max Heals per Ability:":
		return MaxHealsperAbility
	case "Damage beyond 55:":
		return Damagebeyond55
	case "Magical Protection / Debuff:":
		return MagicalProtectionDebuff
	case "Shadow Slow:":
		return ShadowSlow
	case "Light CC Immunity Duration:":
		return LightCCImmunityDuration
	case "Ox Damage:":
		return OxDamage
	case "Attack Speed (Bow):":
		return AttackSpeedBow
	case "Movement Speed":
		return MovementSpeed
	case "Path Duration:":
		return PathDuration
	case "Damage per Wraith:":
		return DamageperWraith
	case "Slow Amount:":
		return SlowAmount
	case "Light Self Heal: ":
		return LightSelfHeal
	case "Physical Penetration:":
		return PhysicalPenetration
	case "Damage Increase:":
		return DamageIncrease
	case "Colossal Duration Increase:":
		return ColossalDurationIncrease
	case "Damage per Dart:":
		return DamageperDart
	default:
		return NoRankItemDescription
	}
}

// MenuItemPart - parsed data regarding a specific part of an ability
type MenuItemPart struct {
	Description string
	Value       string
}

// RankItemPart - parsed data regarding a specific part of an ability
type RankItemPart struct {
	Description string
	Value       string
}

// ItemDescription - parsed data regarding a specific item (ability, basic attack, etc.)
type ItemDescription struct {
	Cooldown             int
	Cost                 []int
	Description          string
	Menuitems            []MenuItemPart
	Rankitems            []RankItemPart
	SecondaryDescription string
}

// Ability - parsed data regarding an ability
type Ability struct {
	AbilityName string
	Description ItemDescription
	ID          int
	Summary     string
	URL         string
}

// God - Holds data pertaining to parsed god
type God struct {
	// Ability structs
	Abilities []Ability
	// God stats
	AttackSpeed                float64
	AttackSpeedPerLevel        float64
	Cons                       string
	HP5PerLevel                float64
	Health                     int
	HealthPerFive              int
	HealthPerLevel             int
	Lore                       string
	MP5PerLevel                float64
	MagicProtection            int
	MagicProtectionPerLevel    float64
	MagicalPower               int
	MagicalPowerPerLevel       int
	Mana                       int
	ManaPerFive                float64
	ManaPerLevel               int
	Name                       string
	OnFreeRotation             string
	Pantheon                   string
	PhysicalPower              int
	PhysicalPowerPerLevel      int
	PhysicalProtection         int
	PhysicalProtectionPerLevel int
	Pros                       string
	Roles                      string
	Speed                      int
	Title                      string
	Type                       string
	// God basic attack
	BasicAttack Ability
	// God image urls
	GodAbilityURLs []string
	GodCardURL     string
	GodIconURL     string
	// God id
	ID int
	// Misc data
	LatestGod bool
	RetMsg    string
}

// rawItemPart - unparsed data regarding a specific part of an ability
type rawItemPart struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// rawItemDescription - unparsed data regarding a specific item (ability, basic attack, etc.)
type rawItemDescription struct {
	Cooldown             string        `json:"cooldown"`
	Cost                 string        `json:"cost"`
	Description          string        `json:"description"`
	Menuitems            []rawItemPart `json:"menuitems"`
	Rankitems            []rawItemPart `json:"rankitems"`
	SecondaryDescription string        `json:"secondaryDescription"`
}

// rawDescription - unparsed data regarding wrapping a rawItemDescription
type rawDescription struct {
	ItemDescription rawItemDescription `json:"itemDescription"`
}

// rawAbility - unparsed data regarding an ability
type rawAbility struct {
	Description rawDescription `json:"Description"`
	ID          int            `json:"Id"`
	Summary     string         `json:"Summary"`
	URL         string         `json:"URL"`
}

// rawGod - Holds data pertaining to unparsed god, intermediate format
type rawGod struct {
	// Ability names
	AbilityName1 string `json:"Ability1"`
	AbilityName2 string `json:"Ability2"`
	AbilityName3 string `json:"Ability3"`
	AbilityName4 string `json:"Ability4"`
	AbilityName5 string `json:"Ability5"`
	// Ability ids
	AbilityID1 int `json:"AbilityId1"`
	AbilityID2 int `json:"AbilityId2"`
	AbilityID3 int `json:"AbilityId3"`
	AbilityID4 int `json:"AbilityId4"`
	AbilityID5 int `json:"AbilityId5"`
	// Ability structs
	AbilityDesciption1 rawAbility `json:"Ability_1"`
	AbilityDesciption2 rawAbility `json:"Ability_2"`
	AbilityDesciption3 rawAbility `json:"Ability_3"`
	AbilityDesciption4 rawAbility `json:"Ability_4"`
	AbilityDesciption5 rawAbility `json:"Ability_5"`
	// God stats
	AttackSpeed                float64 `json:"AttackSpeed"`
	AttackSpeedPerLevel        float64 `json:"AttackSpeedPerLevel"`
	Cons                       string  `json:"Cons"`
	HP5PerLevel                float64 `json:"HP5PerLevel"`
	Health                     int     `json:"Health"`
	HealthPerFive              int     `json:"HealthPerFive"`
	HealthPerLevel             int     `json:"HealthPerLevel"`
	Lore                       string  `json:"Lore"`
	MP5PerLevel                float64 `json:"MP5PerLevel"`
	MagicProtection            int     `json:"MagicProtection"`
	MagicProtectionPerLevel    float64 `json:"MagicProtectionPerLevel"`
	MagicalPower               int     `json:"MagicalPower"`
	MagicalPowerPerLevel       int     `json:"MagicalPowerPerLevel"`
	Mana                       int     `json:"Mana"`
	ManaPerFive                float64 `json:"ManaPerFive"`
	ManaPerLevel               int     `json:"ManaPerLevel"`
	Name                       string  `json:"Name"`
	OnFreeRotation             string  `json:"OnFreeRotation"`
	Pantheon                   string  `json:"Pantheon"`
	PhysicalPower              int     `json:"PhysicalPower"`
	PhysicalPowerPerLevel      int     `json:"PhysicalPowerPerLevel"`
	PhysicalProtection         int     `json:"PhysicalProtection"`
	PhysicalProtectionPerLevel int     `json:"PhysicalProtectionPerLevel"`
	Pros                       string  `json:"Pros"`
	Roles                      string  `json:"Roles"`
	Speed                      int     `json:"Speed"`
	Title                      string  `json:"Title"`
	Type                       string  `json:"Type"`
	// Omitted abilityDescription* fields because they seemed to be duplicates of previous ability fields
	// God basic attack
	BasicAttack rawAbility `json:"basicAttack"`
	// God image urls
	GodAbility1URL string `json:"godAbility1_URL"`
	GodAbility2URL string `json:"godAbility2_URL"`
	GodAbility3URL string `json:"godAbility3_URL"`
	GodAbility4URL string `json:"godAbility4_URL"`
	GodAbility5URL string `json:"godAbility5_URL"`
	GodCardURL     string `json:"godCard_URL"`
	GodIconURL     string `json:"godIcon_URL"`
	// God id
	ID int `json:"id"`
	// Misc data
	LatestGod string `json:"latestGod"`
	RetMsg    string `json:"ret_msg"`
}

func GetRawGods() *[]rawGod {
	file, e := ioutil.ReadFile(G.GodsPath)
	G.ErrorHandler(e)
	rgs := make([]rawGod, 0)
	json.Unmarshal(file, &rgs)
	return &rgs
}
