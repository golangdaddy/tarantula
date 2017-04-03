package validation

import 	(
		"strings"
		)

const 	(
		COUNTRY_CSV = `POPULATED,COUNTRY,A2 (ISO),A3 (UN),REGION,LATITUDE,LONGITUDE,OFFICIAL LANGUAGE,nada,RECOGNISED LANGUAGES,ISO 639-1 CODES,NUM (UN),DIALING CODE,COMMENT,NUMBER OF PHOTOS FOR FRONTPAGE,TRIGGERING WHEN SWITCHING COUNTRY,TRIGGERING BASED ON IP
,United Kingdom ,GB,GBR,Europe,55.378051,-3.435973,EN,en,Scots, Ulster Scots, Welsh, Cornish, Scottish Gaelic, Irish,,826,44,,3,,
,Denmark,DK,DNK,Europe,56.263920,9.501785,DA,da,Faroese, Greenlandic, German,,208,45,,3,YES,YES
,Spain,ES,ESP,Europe,40.463667,-3.749220,ES,es,Basque, Catalan, Galician,,724,34,,3,,
,Isle of Man,IM,IMN,Europe,54.236107,-4.548056,EN,en,Manx,,833,44,,1,,
,United Arab Emirates,AE,ARE,Middle East,23.424076,53.847818,AR,,,,784,971,,,,
,Jersey,JE,JEY,Europe,49.214439,-2.131250,EN,en,French, Jèrriais,,832,44,,,,
,Guernsey,GG,GGY,Europe,49.465691,-2.585278,EN,en,Guernésiais, Sercquiais, Auregnais,,831,44,,,,
,Faroe Islands,FO,FRO,Europe,61.892635,-6.911806,DA,fo,Danish,,234,298,,,,
,Greenland,GL,GRL,Europe,71.706936,-42.604303,KL,kl,Danish,,304,299,,,,
,Malta,MT,MLT,Europe,35.937496,14.375416,MT,mt,English, Maltese Sign Language,,470,356,,,,
,Liechtenstein,LI,LIE,Europe,47.166000,9.555373,DE,de,,,438,423,,,,
,Luxembourg,LU,LUX,Europe,49.815273,6.129583,FR,fr,Luxembourgish, German,,442,352,,,,
,Andorra,AD,AND,Europe,42.546245,1.601554,CA,ca,Spanish, French, Portuguese,,20,376,,,,
,Antigua and Barbuda,AG,ATG,Caribbean,17.060816,-61.796428,EN,en,Antiguan Creole,,28,1-268,,,,
,Monaco,MC,MCO,Europe,43.750298,7.412841,FR,fr,Monégasque, Italian, Occitan, English,,492,377,,,,
,Ireland,IE,IRL,Europe,53.412910,-8.243890,EN,en,Irish, Ulster Scots,,372,353,,,,
,Switzerland,CH,CHE,Europe,46.818188,8.227512,DE,de,French, Italian, Romansh,,756,41,,,,
,Gibraltar,GI,GIB,Europe,36.137741,-5.345374,ES,en,Llanito,,292,350,,,,
,Iceland,IS,ISL,Europe,64.963051,-19.020835,IS,is,,,352,354,,,,
,Bahamas,BS,BHS,Caribbean,25.034280,-77.396280,EN,en,,,44,1-242,,,,
,Barbados,BB,BRB,Caribbean,13.193887,-59.543198,EN,en,Barbadian (Bajan Creole),,52,1-246,,,,
,Belize,BZ,BLZ,Central America,17.189877,-88.497650,EN,en,,,84,501,,,,
,Dominica,DM,DMA,Caribbean,15.414999,-61.370976,EN,en,Kwéyòl, Island Carib,,212,1-767,,,,
,Saint Kitts and Nevis,KN,KNA,Caribbean,17.357822,-62.782998,EN,en,,,659,1-869,,,,
,Jamaica,JM,JAM,Caribbean,18.109581,-77.297508,EN,en,Jamaican Patois,,388,1-876,,,,
,Saint Vincent and the Grenadines,VC,VCT,Caribbean,12.984305,-61.287228,EN,en,,,670,1-784,,,,
,British Virgin Islands,VG,VGB,Caribbean,18.420695,-64.639968,EN,en,,,92,1-284,,,,
,US Virgin Islands,VI,VIR,Caribbean,18.335765,-64.896335,EN,en,Spanish, French,,850,1-340,,,,
,Bermuda,BM,BMU,Caribbean,32.321384,-64.757370,EN,en,,,60,1-441,,,,
,Cayman Islands,KY,CYM,Caribbean,19.513469,-80.566956,EN,en,,,136,1-345,,,,
,Trinidad and Tobago,TT,TTO,Caribbean,10.691803,-61.222503,EN,en,Trinidadian Creole,,780,1-868,,,,
,Grenada,GD,GRD,Caribbean,12.262776,-61.604171,EN,en,Grenadian Creole English, Grenadian Creole French,,308,1-473,,,,
,Sweden,SE,SWE,Europe,60.128161,18.643501,SV,sv,,,752,46,towns > 15,000,,,
,Norway,NO,NOR,Europe,60.472024,8.468946,NO,no,Kven, Romani, Scandoromani, Northern Sami, Lule Sami, Southern Sami,,578,47,towns > 7,000,,,
,Finland,FI,FIN,Europe,61.924110,25.748151,FI,fi,Swedish, Sami,,246,358,,,,
,Germany,DE,DEU,Europe,51.165691,10.451526,DE,de,,,276,49,,3,,
,Estonia,EE,EST,Europe,58.595272,25.013607,ET,et,Võro, Setu,,233,372,,,,
,Latvia,LV,LVA,Europe,56.879635,24.603189,LV,,,,428,371,,,,
,Lithuania,LT,LTU,Europe,55.169438,23.881275,LT,,,,440,370,,,,
,Thailand,TH,THA,Southeast Asia ,15.870032,100.992541,TH,,,,764,66,,,,
,France,FR,FRA,Europe,46.227638,2.213749,FR,fr,,,250,33,,,,
,Italy,IT,ITA,Europe,41.871940,12.567380,IT,,,,380,39,,,,
,Israel,IL,ISR,Middle-East,31.046051,34.851612,IW,,Arabic,,376,972,,,,
,Netherlands,NL,NLD,Europe,52.132633,5.291266,NL,,English, West Frisian, Papiamento, Limburgish, Dutch Low Saxon,,528,31,,,,
,Belgium,BE,BEL,Europe,50.503887,4.469936,FR,fr,Dutch, German,,56,32,,,,
,Greece,GR,GRC,Europe,39.074208,21.824312,EL,,,,300,30,,,,
,Turkey,TR,TUR,Middle East,38.963745,35.243322,TR,,,,792,90,,,,
,Saint Lucia,LC,LCA,Caribbean,13.909444,-60.978893,EN,en,Saint Lucian Creole, French,,662,1-758,,,,
,Anguilla,AI,AIA,Caribbean,18.220554,-63.068615,EN,en,,,660,1-264,,,,
,Curaçao,CW,CUW,Caribbean,12.116667,-68.933333,NL,,Papiamentu, English,,531,599,,,,
,Puerto Rico,PR,PRI,Caribbean,18.220833,-66.590149,ES,es,English,,630,1,,,,
,Saint Martin (French part),MF,MAF,Caribbean,18.066667,-63.050000,FR,fr,English,,663,590,,,,
,Sint Maarten (Dutch part),SX,SXM,Caribbean,18.033333,-63.050000,NL,,English,,534,1-721,,,,
,Turks and Caicos Islands,TC,TCA,Caribbean,21.694025,-71.797928,EN,en,,,796,1-649,,,,
,Aruba,AW,ABW,Caribbean,12.521110,-69.968338,NL,,Papiemento,,533,297,,,,
,Bonaire,BQ,BES,Caribbean,12.150000,-68.266667,NL,,Papiemento,,535,599,,,,
,Cuba,CU,CUB,Caribbean,21.521757,-77.781167,ES,es,,,192,53,,,,
,Dominican Republic,DO,DOM,Caribbean,18.735693,-70.162651,ES,es,,,214,1-809,1-829,1-849,,,,
,Martinique,MQ,MTQ,Caribbean,14.641528,-61.024174,FR,fr,Martiniquan Creole,,474,596,,,,
,Montserrat,MS,MSR,Caribbean,16.742498,-62.187366,EN,en,,,500,1-664,,,,
,Guadeloupe,GP,GLP,Caribbean,16.995971,-62.067641,FR,fr,Creole,,312,590,,,,
,Saint Barthalemy,BL,BLM,Caribbean,17.900000,-62.833333,FR,fr,Saint-Barthélemy French, Antillean Creole,,652,590,,,,
,Haiti,HT,HTI,Caribbean,18.971187,-72.285215,FR,fr,Haitian Creole,,332,509,,,,
,Cook Islands,CK,COK,Oceania,-21.236736,-159.777671,EN,en,Cook Islands Māori, Pukapukan, Rakahanga-Manihiki,,184,682,,,,
,Micronesia,FM,FSM,Oceania,7.425554,150.550812,EN,en,Chuukese, Kosraen, Woleaian,,583,691,,,,
,Guyana,GY,GUY,South America,4.860416,-58.930180,EN,en,Portuguese, Spanish, Akawaio, Macushi, Waiwai, Arawak, Patamona, Warrau, Carib, Wapishana, Arekuna,,328,592,,,,
,Kiribati,KI,KIR,Oceania,-3.370417,-168.734039,EN,en,Gilbertese,,296,686,,,,
,Marshall Islands,MH,MHL,Oceania,7.131474,171.184478,EN,,English,,584,692,,,,
,Nauru,NR,NRU,Oceania,-0.522778,166.931503,NA,,English,,520,674,,,,
,Niue,NU,NIU,Oceania,-19.054445,-169.867233,EN,,English,,570,683,,,,
,Palau,PW,PLW,Oceania,7.514980,134.582520,EN,,English, Japanese, Sonsorolese, Tobian,,585,680,,,,
,Papua New Guinea,PG,PNG,Oceania,-6.314993,143.955550,EN,,Tok Pisin, Papua New Guinean Sign Language, English,,598,675,,,,
,Samoa,WS,WSM,Oceania,-13.759029,-172.104629,SM,,English,,882,685,,,,
,Seychelles,SC,SYC,Africa,-4.679574,55.491977,EN,en,French, Seychellois Creole,,690,248,,,,
,Solomon Islands,SB,SLB,Oceania,-9.645710,160.156194,EN,en,,,90,677,,,,
,Tonga,TO,TON,Oceania,-21.178986,-175.198242,TO,,English,,776,676,,,,
,Tuvalu,TV,TUV,Oceania,-7.109535,177.649330,EN,,English,,798,688,,,,
,Vanuatu,VU,VUT,Oceania,-15.376706,166.959158,BI,,French, English,,548,678,,,,
,American Samoa,AS,ASM,Oceania,-14.270972,-170.132217,EN,,English,,16,1-684,,,,
,Christmas Island,CX,CXR,Oceania,-10.447525,105.690449,MS,,Chinese, English,,162,61,,,,
,Falkland Islands (Malvinas),FK,FLK,South America,-51.796253,-59.523613,EN,en,spanish,,238,500,,,,
,Guam,GU,GUM,Oceania,13.444304,144.793731,EN,en,Chamorro,,316,1-671,,,,
,Norfolk Island,NF,NFK,Oceania,-29.040835,167.954712,EN,,English,,574,672,,,,
,Northern Mariana Islands,MP,MNP,Oceania,17.330830,145.384690,EN,en,Chamorro, Carolinian,,580,1-670,,,,
,Pitcairn,PN,PCN,Oceania,-24.703615,-127.439308,EN,,English,,612,870,,,,
,New Zealand,NZ,NZL,Oceania,-40.900557,174.885971,EN,en,Māori, New Zealand Sign Language,,554,64,,,,
,Australia,AU,AUS,Oceania,-25.274398,133.775136,EN,en,,,36,61,,,,
,Afghanistan,AF,AFG,Asia,33.939110,67.709953,PS,,Dari,,4,93,,,,
,Albania,AL,ALB,Europe,41.153332,20.168331,SQ,,,,8,355,,,,
,Algeria,DZ,DZA,Africa,28.033886,1.659626,AR,,Berber,,12,213,,,,
,Angola,AO,AGO,Africa,-11.202692,17.873887,PT,,Kikongo, Chokwe, Umbundu, Kimbundu, Nganguela, Kwanyama,,24,244,,,,
,Argentina,AR,ARG,South America,-38.416097,-63.616672,ES,es,,,32,54,,,,
,Armenia,AM,ARM,Russia and CIS,40.069099,45.038189,HY,,Eastern Armenian, Armenian alphabet,,51,374,,,,
,Austria,AT,AUT,Europe,47.516231,14.550072,DE,de,,,40,43,,,,
,Azerbaijan,AZ,AZE,Russia and CIS,40.143105,47.576927,AZ,,,,31,994,,,,
,Bahrain,BH,BHR,Middle East,25.930414,50.637772,AR,,,,48,973,,,,
,Bangladesh,BD,BGD,Asia,23.684994,90.356331,BN,,English,,50,880,,,,
,Belarus,BY,BLR,Russia and CIS,53.709807,27.953389,BE,,Russian,,112,375,,,,
,Benin,BJ,BEN,Africa,9.307690,2.315834,FR,fr,Fon, Yoruba,,204,229,,,,
,Bhutan,BT,BTN,Asia,27.514162,90.433601,EN,,,,64,975,,,,
,Bolivia,BO,BOL,South America,-16.290154,-63.588653,ES,es,Quechua, Aymara, Guarani,,68,591,,,,
,Bosnia and Herzegovina,BA,BIH,Europe,43.915886,17.679076,HR,,Croatian, Serbian,,70,387,,,,
,Botswana,BW,BWA,Africa,-22.328474,24.684866,TN,,English,,72,267,,,,
,Brazil,BR,BRA,South America,-14.235004,-51.925280,PT,,,,76,55,,,,
,British Indian Ocean Territory,IO,IOT,Asia,-6.343194,71.876519,EN,en,,,86,246,,,,
,Brunei Darussalam,BN,BRN,Southeast Asia ,4.535277,114.727669,MS,,English, Brunei Malay, Jawi Malay, Tutong dialect, Kedayan dialect, Belait dialect, Chinese, Murut, Dusun, Brunei Bisaya,,96,673,,,,
,Bulgaria,BG,BGR,Europe,42.733883,25.485830,BG,,,,100,359,,,,
,Burkina Faso,BF,BFA,Africa,12.238333,-1.561593,FR,fr,Mòoré, Mandinka, Bambara,,854,226,,,,
,Burundi,BI,BDI,Africa,-3.373056,29.918886,FR,fr,Kirundi,,108,257,,,,
,Cambodia,KH,KHM,Southeast Asia ,12.565679,104.990963,EN,,,,116,855,,,,
,Cameroon,CM,CMR,Africa,7.369722,12.354722,FR,fr,English,,120,237,,,,
,Canada,CA,CAN,North America,56.130366,-106.346771,EN,en,French, Chipewyan, Cree, Gwich'in, Inuinnaqtun, Inuktitut, Inuvialuktun, North Slavey, South Slavey, Tłı̨chǫ,,124,1,,,,
,Cape Verde,CV,CPV,Africa,16.002082,-24.013197,PT,,Cape Verdean Creole,,132,238,,,,
,Central African Republic,CF,CAF,Africa,6.611111,20.939444,SG,,French,,140,236,,,,
,Chad,TD,TCD,Africa,15.454166,18.732207,AR,,French,,148,235,,,,
,Chile,CL,CHL,South America,-35.675147,-71.542969,ES,es,,,152,56,,,,
,China,CN,CHN,Asia,35.861660,104.195397,ZH,,Mongolian, Tibetan, Uyghur, Zhuang,,156,86,,,,
,Cocos (Keeling) Islands,CC,CCK,Oceania,-12.164165,96.870956,MS,,English,,166,61,,,,
,Colombia,CO,COL,South America,4.570868,-74.297333,ES,es,English,,170,57,,,,
,Comoros,KM,COM,Africa,-11.875001,43.872219,FR,,Arabic, French,,174,269,,,,
,Republic of the Congo,CG,COG,Africa,-0.228021,15.827659,FR,fr,Kituba, Lingala,,178,242,,,,
,Democratic Republic of the Congo,CD,COD,Africa,-4.038333,21.758664,FR,fr,Lingala, Kituba, Swahili, Tshiluba,,180,243,,,,
,Costa Rica,CR,CRI,Central America,9.748917,-83.753428,ES,es,Mekatelyu, Bribri, Patois,,188,506,,,,
,Croatia,HR,HRV,Europe,45.100000,15.200000,HR,,,,191,385,,,,
,Cyprus,CY,CYP,Middle East,35.126413,33.429859,EL,,Turkish, Armenian, Cypriot Arabic, Cypriot Greek, Cypriot Turkish,,196,357,,,,
,Czechia,CZ,CZE,Europe,49.817492,15.472962,CS,,Slovak, German, Polish, Belarusian, Bulgarian, Croatian, Greek, Hungarian, Romani, Russian, Rusyn, Serbian, Ukranian, Vietnamese,,203,420,,,,
,Czech Republic,CZ,CZE,Europe,49.817492,15.472962,CS,,Slovak, German, Polish, Belarusian, Bulgarian, Croatian, Greek, Hungarian, Romani, Russian, Rusyn, Serbian, Ukranian, Vietnamese,,203,420,,,,
,Côte d'Ivoire,CI,CIV,Africa,7.539989,-5.547080,FR,fr,Bété, Dioula, Baoulé, Abron, Agni, Cebaara, Senufo,,384,225,,,,
,Djibouti,DJ,DJI,Africa,11.825138,42.590275,FR,fr,Arabic, Somali, Afar,,262,253,,,,
,Ecuador,EC,ECU,South America,-1.831239,-78.183406,ES,es,Kichwa, Shuar,,218,593,,,,
,Egypt,EG,EGY,Middle East,26.820553,30.802498,AR,,,,818,20,,,,
,El Salvador,SV,SLV,Central America,13.794185,-88.896530,ES,es,,,222,503,,,,
,Equatorial Guinea,GQ,GNQ,Africa,1.650801,10.267895,ES,es,French, Portuguese, Fang, Bube, Igbo, Pidgin English, Annobonese,,226,240,,,,
,Eritrea,ER,ERI,Africa,15.179384,39.782334,TI,,Tigre, Kunama, Saho, Bilen, Nara, Afar,,232,291,,,,
,Ethiopia,ET,ETH,Africa,9.145000,40.489673,AM,,,,231,251,,,,
,Fiji,FJ,FJI,Oceania,-16.578193,179.414413,FJ,,English, Hindi,,242,679,,,,
,French Guiana,GF,GUF,South America,3.933889,-53.125782,FR,fr,French Guianese Creole, Arawak, Palijur, Kali'na, Wayana, Wayampi, Emirillon, Saramaka, Paramaccan, Aluku, Ndyuka, Hmong Njua, Portuguese, Hakka, Haitian Creole, Spanish, Dutch, English,,254,594,,,,
,French Polynesia,PF,PYF,Oceania,-17.679742,-149.406843,FR,fr,,,258,689,,,,
,French Southern Territories,TF,ATF,Africa,-49.280366,69.348557,FR,fr,,,260,262,,,,
,Gabon,GA,GAB,Africa,-0.803689,11.609444,FR,fr,Fang, Myene, Punu, Nzebi,,266,241,,,,
,Gambia,GM,GMB,Africa,13.443182,-15.310139,EN,en,Mandinka, Fula, Wolof, Serer, Jola,,270,220,,,,
,Georgia,GE,GEO,Asia,42.315407,43.356892,KA,,Russian, Armenian, Azerbaijani,,268,995,,,,
,Ghana,GH,GHA,Africa,7.946527,-1.023194,EN,en,Akuapem Twi, Asante Twi, Dagaare, Dagbani, Dangme, Ewe, Ga, Gonja, Kasem, Fante, Nzema, Wasa, Talensi, Frafra,,288,233,,,,
,Guatemala,GT,GTM,Central America,15.783471,-90.230759,ES,en,,,320,502,,,,
,Guinea,GN,GIN,Africa,9.945587,-9.696645,FR,fr,Maninka Fula, Susu,,324,224,,,,
,Guinea-Bissau,GW,GNB,Africa,11.803749,-15.180413,PT,,Upper Guinea Creole,,624,245,,,,
,Haiti,HT,HTI,Caribbean,18.971187,-72.285215,FR,fr,Haitian Creole,,332,509,,,,
,Holy See (Vatican City State),VA,VAT,Europe,41.902916,12.453389,IT,,inga,,336,379,,,,
,Honduras,HN,HND,Central America,15.199999,-86.241905,ES,es,,,340,504,,,,
,Hong Kong,HK,HKG,Asia,22.396428,114.109497,EN,en,Chinese, Cantonese,,344,852,,,,
,Hungary,HU,HUN,Europe,47.162494,19.503304,HU,,,,348,36,,,,
,India,IN,IND,Asia,20.593684,78.962880,HI,,English, Assamese, Bengali, Bodo, Dogri, Gujarati, Kannada, Kashmiri, Konkani, Maithili, Malayalam, Manipuri, Marathi, Nepali, Odia, Punjabi, Sanskrit, Santali, Sindhi, Tamil, Telugu, Urdu,,356,91,,,,
,Indonesia,ID,IDN,Southeast Asia ,-0.789275,113.921327,IN,,,,360,62,,,,
,Iran,IR,IRN,Middle East,32.427908,53.688046,FA,,Azerbaijani, Kurdish, Lurish, Samnani, Gilaki, Mazandarani, Tati, Turkmen, Arabic, Baloch, Talysh, Georgian, Armenian, Neo-Aramaic,,364,98,,,,
,Iraq,IQ,IRQ,Middle East,33.223191,43.679291,AR,,Kurdish,,368,964,,,,
,Japan,JP,JPN,Asia,36.204824,138.252924,JA,,Aynu Itak, Amami, Kikai, Kunigami, Miyako, Okinawan, Okinoerabu, Tokunoshima, Yaeyama, Yonaguni, Yoron,,392,81,,,,
,Jordan,JO,JOR,Middle East,30.585164,36.238414,AR,,,,400,962,,,,
,Kazakhstan,KZ,KAZ,Russia and CIS,48.019573,66.923684,KK,,Russian,,398,7,,,,
,Kenya,KE,KEN,Africa,-0.023559,37.906193,SW,,English,,404,254,,,,
,North Korea,KP,PRK,Asia,40.339852,127.510093,KO,,,,408,850,,,,
,South Korea,KR,KOR,Asia,35.907757,127.766922,KO,,,,410,82,,,,
,Kuwait,KW,KWT,Middle East,29.311660,47.481766,AR,,,,414,965,,,,
,Kyrgyzstan,KG,KGZ,Russia and CIS,41.204380,74.766098,RU,,Russian,,417,996,,,,
,Lao People's Democratic Republic,LA,LAO,Southeast Asia ,19.856270,102.495496,LO,,French, Hmong, Khmu,,418,856,,,,
,Lebanon,LB,LBN,Middle East,33.854721,35.862285,AR,,,,422,961,,,,
,Lesotho,LS,LSO,Africa,-29.609988,28.233608,ST,,English,,426,266,,,,
,Liberia,LR,LBR,Africa,6.428055,-9.429499,EN,en,Liberian English,,430,231,,,,
,Libya,LY,LBY,Africa,26.335100,17.228331,AR,,Libyan Arabic, Tamazight, Italian,,434,218,,,,
,Macao,MO,MAC,Asia,22.198745,113.543873,ZH,,Portuguese, Cantonese,,446,853,,,,
,Macedonia (FYROM),MK,MKD,Europe,41.608635,21.745275,MK,,Albanian, Turkish, Romani, Serbian,,807,389,,,,
,Madagascar,MG,MDG,Africa,-18.766947,46.869107,MG,,French,,450,261,,,,
,Malawi,MW,MWI,Africa,-13.254308,34.301525,EN,en,Chichewa,,454,265,,,,
,Malaysia,MY,MYS,Southeast Asia ,4.210484,101.975766,MS,,English,,458,60,,,,
,Maldives,MV,MDV,Asia,3.202778,73.22068,EN,,,,462,960,,,,
,Mali,ML,MLI,Africa,17.570692,-3.996166,FR,fr,Bambara,,466,223,,,,
,Mauritania,MR,MRT,Africa,21.007890,-10.940835,AR,,Pulaar, Soninke, Wolof, French, Zenaga Berber,,478,222,,,,
,Mauritius,MU,MUS,Africa,-20.348404,57.552152,EN,en,French, Bhojpuri, Mauritian Creole, Arabic, Chinese, Hindi, Mandarin, Marathi, Sanskrit, Tamil, Telugu, Urdu,,480,230,,,,
,Mayotte,YT,MYT,Africa,-12.827500,45.166244,FR,,Kibushi, Kiantalaotsi, French, Arabic,,175,262,,,,
,Mexico,MX,MEX,North America,23.634501,-102.552784,ES,es,Nahuatl, Yucatec Maya, Zapotec, Mixtec, Mayo, Yaqui, Tzeltal, Tzotzil, Chol, Totonac, Purépecha, Otomi, Mazahua, Mazatec, Chinantec, Mixe, Zoque, Populuca, Popoloca language, Me'phaa, Wizarika, Naayerite, Tepehuán, Warihio, Raramuri, Seri, Chontal Maya, Chontal, Huave, Pame, Teenek, Kickapoo, Kiliwa, Paipai, Cucapá, Amuzgo, Triqui, Lacandon Maya, Mam Maya, Jakaltek, Matlatzinca, Tepehua, Chichimeca Jonaz, Pima Bjao, Ngiwa, Ixcatec, Ayapanec, Catalan, Plautdietsch, Chipilo, English, German, Greek, Italian, Arabic, French, Portuguese, Chinese, Japanese, Mexican Sign Language, Yucatan Sign Language, American Sign Language,,484,52,,,,
,Moldova,MD,MDA,Russia and CIS,47.411631,28.369885,RO,,,,498,373,,,,
,Mongolia,MN,MNG,Asia,46.862496,103.846656,MN,,,,496,976,,,,
,Montenegro,ME,MNE,Europe,42.708678,19.374390,EN,,Serbian, Bosnian, Albanian, Croatian,,499,382,,,,
,Morocco,MA,MAR,Africa,31.791702,-7.092620,AR,,Berber, French, Moroccan Arabic, Hassaniya Arabic,,504,212,,,,
,Mozambique,MZ,MOZ,Africa,-18.665695,35.529562,PT,,,,508,258,,,,
,Myanmar,MM,MMR,Southeast Asia ,21.913965,95.956223,MY,,Jingpho, Kayah, Karen, Chin, Mon, Rakhine, Shan,,104,95,,,,
,Namibia,NA,NAM,Africa,-22.957640,18.490410,EN,en,Afrikaans, German, Ju'hoansi, Khoekhoegowab, Oshikwanyama, Oshindonga, Otjiherero, Rukwangali, Rumanyo, Setswana, Silozi, Thimbukushu,,516,264,,,,
,Nepal,NP,NPL,Asia,28.394857,84.124008,NE,,,,524,977,,,,
,New Caledonia,NC,NCL,Oceania,-20.904305,165.618042,FR,,,,540,687,,,,
,Nicaragua,NI,NIC,Central America,12.865416,-85.207229,EN,,,,558,505,,,,
,Niger,NE,NER,Africa,17.607789,8.081666,FR,,,,562,227,,,,
,Nigeria,NG,NGA,Africa,9.081999,8.675277,YO,,,,566,234,,,,
,Oman,OM,OMN,Middle East,21.512583,55.923255,AR,,,,512,968,,,,
,Pakistan,PK,PAK,Asia,30.375321,69.345116,UR,,,,586,92,,,,
,Palestinian Territories,PS,PSE,Middle East,31.952162,35.233154,AR,,,,275,970,,,,
,Panama,PA,PAN,Central America,8.537981,-80.782127,ES,,,,591,507,,,,
,Paraguay,PY,PRY,South America,-23.442503,-58.443832,ES,,,,600,595,,,,
,Peru,PE,PER,South America,-9.189967,-75.015152,ES,,,,604,51,,,,
,Philippines,PH,PHL,Southeast Asia ,12.879721,121.774017,EN,,,,608,63,,,,
,Poland,PL,POL,Europe,51.919438,19.145136,PL,,,,616,48,,,,
,Portugal,PT,PRT,Europe,39.399872,-8.224454,PT,,,,620,351,,,,
,Qatar,QA,QAT,Middle East,25.354826,51.183884,AR,,,,634,974,,,,
,Romania,RO,ROU,Europe,45.943161,24.966760,RO,,,,642,40,,,,
,Russian Federation,RU,RUS,Russia and CIS,61.524010,105.318756,RU,,,,643,7,,,,
,Rwanda,RW,RWA,Africa,-1.940278,29.873888,FR,,,,646,250,,,,
,Reunion,RE,REU,Africa,-21.115141,55.536384,FR,,,,638,262,,,,
,Saint Helena,SH,SHN,Africa,-24.143474,-10.030696,EN,,,,654,290,,,,
,Saint Pierre and Miquelon,PM,SPM,North America,46.941936,-56.271110,EN,,,,666,508,,,,
,San Marino,SM,SMR,Europe,43.942360,12.457777,EN,,,,674,378,,,,
,Sao Tome and Principe,ST,STP,Africa,0.186360,6.613081,EN,,,,678,239,,,,
,Saudi Arabia,SA,SAU,Middle East,23.885942,45.079162,AR,,,,682,966,,,,
,Senegal,SN,SEN,Africa,14.497401,-14.452362,FR,,,,686,221,,,,
,Serbia,RS,SRB,Europe,44.016521,21.005859,SR,,,,688,381,,,,
,Sierra Leone,SL,SLE,Africa,8.460555,-11.779889,EN,,,,694,232,,,,
,Singapore,SG,SGP,Southeast Asia ,1.352083,103.819836,MA,,,,702,65,,,,
,Slovakia,SK,SVK,Europe,48.669026,19.699024,SK,,,,703,421,,,,
,Slovenia,SI,SVN,Europe,8.460555,-11.779889,SL,,,,705,386,,,,
,Solomon Islands,SB,SLB,Oceania,-9.645710,160.156194,EN,,,,90,677,,,,
,Somalia,SO,SOM,Africa,5.152149,46.199616,AR,,,,706,252,,,,
,South Africa,ZA,ZAF,Africa,-30.559482,22.937506,EN,,,,710,27,,,,
,South Georgia and the South Sandwich Islands,GS,SGS,South America,-54.429579,-36.587909,EN,,,,239,500,,,,
,South Sudan,SS,SSD,Africa,6.876992,31.306979,EN,,,,728,211,,,,
,Sri Lanka,LK,LKA,Asia,7.873054,80.771797,EN,,,,144,94,,,,
,Sudan,SD,SDN,Africa,12.862807,30.217636,AR,,,,729,249,,,,
,Suriname,SR,SUR,South America,3.919305,-56.027783,NL,,,,740,597,,,,
,Svalbard and Jan Mayen,SJ,SJM,Europe,77.553604,23.670272,NO,,,,744,47,,,,
,Swaziland,SZ,SWZ,Africa,-26.522503,31.465866,SS,,,,748,268,,,,
,Syrian Arab Republic,SY,SYR,Middle East,34.802075,38.996815,AR,,,,760,963,,,,
,Taiwan,TW,TWN,Asia,23.697810,120.960515,ZH,,,,158,886,,,,
,Tajikistan,TJ,TJK,Russia and CIS,38.861034,71.276093,TG,,,,762,992,,,,
,United Republic of Tanzania,TZ,TZA,Africa,-6.369028,34.888822,SW,,,,834,255,,,,
,Timor-Leste,TL,TLS,Southeast Asia ,-8.874217,125.727539,PT,,,,626,670,,,,
,Togo,TG,TGO,Africa,8.619543,0.824782,FR,,,,768,228,,,,
,Tokelau,TK,TKL,Oceania,-8.967363,-171.855881,EN,,,,772,690,,,,
,Tunisia,TN,TUN,Africa,33.886917,9.537499,AR,,,,788,216,,,,
,Turkmenistan,TM,TKM,Russia and CIS,38.969719,59.556278,TR,,,,795,993,,,,
,Uganda,UG,UGA,Africa,1.373333,32.290275,EN,,,,800,256,,,,
,Ukraine,UA,UKR,Russia and CIS,48.379433,31.165580,UK,,,,804,380,,,,
,United States,US,USA,North America,37.090240,-95.712891,EN,,,,840,1,,,,
,United States Minor Outlying Islands,UM,UMI,North America,,,EN,,,,581,1,,,,
,Uruguay,UY,URY,South America,-32.522779,-55.765835,ES,,,,858,598,,,,
,Uzbekistan,UZ,UZB,Russia and CIS,41.377491,64.585262,UZ,,,,860,998,,,,
,Venezuela,VE,VEN,South America,6.423750,-66.58973,ES,,,,862,58,,,,
,Viet Nam,VN,VNM,Southeast Asia ,14.058324,108.277199,VI,,,,704,84,,,,
,Wallis and Futuna,WF,WLF,Oceania,-13.768752,-177.156097,FR,,,,876,681,,,,
,Western Sahara,EH,ESH,Africa,24.215527,-12.885834,AR,,,,732,212,,,,
,Yemen,YE,YEM,Middle East,15.552727,48.516388,AR,,,,887,967,,,,
,Zambia,ZM,ZMB,Africa,-13.133897,27.849332,EN,,,,894,260,,,,
,Zimbabwe,ZW,ZWE,Africa,-19.015438,29.154857,EN,,,,716,263,,,,
,Aland Islands,AX,ALA,Europe,60.178525,19.915610,SV,,,,248,358,,,,`
		)

type Country struct {
	Name string 				`json:"name"`
	Code string 				`json:"code"`
	Lat string 					`json:"lat"`
	Lng string 					`json:"lng"`
	Lang string 				`json:"lang"`
}

func parseCountry(s string) *Country {

	cell := strings.Split(s, ",")

	name := strings.TrimSpace(cell[1])
	code := strings.TrimSpace(cell[2])
	lat := strings.TrimSpace(cell[5])
	lng := strings.TrimSpace(cell[6])
	lang := strings.TrimSpace(cell[7])

	return &Country{name, code, lat, lng, lang}
}

func Countries() map[string]*Country {

	m := map[string]*Country{}

	for _, line := range strings.Split(COUNTRY_CSV, "\n")[1:] {

		c := parseCountry(line)

		m[c.Code] = c

	}

	return m
}

func ReverseCountries() map[string]*Country {

	m := map[string]*Country{}

	for _, line := range strings.Split(COUNTRY_CSV, "\n")[1:] {

		c := parseCountry(line)

		m[strings.ToLower(c.Name)] = c

	}

	return m
}