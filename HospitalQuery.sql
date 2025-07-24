--Database and table creation

--CREATE Database hospital;
CREATE Table Patients(
	patient_id SERIAL PRIMARY KEY,
	last_name VARCHAR(50) Not Null,
	first_name VARCHAR(50) Not Null,
	age INT Not Null,
	gender VARCHAR(1) Not Null,
	email VARCHAR(50) UNIQUE Not Null,
	phone VARCHAR(10) UNIQUE Not Null,
	address VARCHAR(50) Not Null,
	credit NUMERIC(10,2) Not Null
);
CREATE Table Departments(
	department_id SERIAL PRIMARY KEY,
	name VARCHAR(50) UNIQUE Not Null,
	rooms_number INT Not Null,
	description VARCHAR(450) Not Null
);
CREATE Table Doctors(
	doctor_id SERIAL PRIMARY KEY,
	last_name VARCHAR(50) Not Null,
	first_name VARCHAR(50) Not Null,
	specialty VARCHAR(50) Not Null,
	email VARCHAR(50) UNIQUE Not Null,
	phone VARCHAR(11) UNIQUE Not Null,
	department_id INT Not Null,
	FOREIGN KEY (department_id) References departments(department_id) ON DELETE CASCADE,
	is_Headdoctor BOOLEAN Not Null,
	credit NUMERIC(10,2) Null
);
CREATE Table Medications(
	medication_id SERIAL PRIMARY KEY,
	name_medication VARCHAR(50) UNIQUE Not Null,
	description VARCHAR(455) Not Null,
	price DECIMAL Not Null, 
	department_id INT Not Null,
	FOREIGN KEY (department_id) References Departments(department_id) ON DELETE CASCADE,
	inventory INT Not Null
);
CREATE Table Rooms(
	room_id SERIAL PRIMARY KEY,
	number_room INT UNIQUE Not Null,
	department_id INT Not Null,
	FOREIGN KEY (department_id) References Departments(department_id) ON DELETE CASCADE,
	type_room VARCHAR(50) Not Null,
	capacity INT Not Null
);
CREATE Table Appointments(
	appointment_id SERIAL PRIMARY KEY,
	patient_id INT Not Null,
	doctor_id INT Not Null,
	room_id INT Not Null,
	FOREIGN KEY (patient_id) References Patients(patient_id) ON DELETE CASCADE,
	FOREIGN KEY (doctor_id) References Doctors(doctor_id) ON DELETE CASCADE,
	FOREIGN KEY (room_id) References Rooms(room_id) ON DELETE CASCADE,
	schedule_date TIMESTAMP Not Null,
	status VARCHAR(50) Not Null
);
CREATE Table Examinations(
	examination_id SERIAL PRIMARY KEY,
	appointment_id INT Not Null,
	FOREIGN KEY (appointment_id) References Appointments(appointment_id) ON DELETE CASCADE,
	examination_date TIMESTAMP Not Null,
	information VARCHAR(300) Not Null,
	assistant_nurses TEXT Not Null,
	diagnosis VARCHAR(100) Not Null,
	result VARCHAR(20) Not Null,
	reexamination_date TIMESTAMP Null,
	price DECIMAL Not Null,
	level VARCHAR(10) Not Null,
	examination_notes TEXT Null
);
CREATE Table Prescriptions(
	prescription_id SERIAL PRIMARY KEY,
	examination_id INT Null,
	FOREIGN KEY (examination_id) References Examinations(examination_id) ON DELETE SET NULL,
	number_medications INT Not Null,
	notes TEXT Not Null,
	prescription_date TIMESTAMP Not Null,
	total_price NUMERIC(10,2) Not Null
);
CREATE Table Nurses(
	nurse_id SERIAL PRIMARY KEY,
	last_name VARCHAR(50) Not Null,
	first_name VARCHAR(50) Not Null,
	phone VARCHAR(10) UNIQUE Not Null,
	department_id INT Not Null,
	FOREIGN KEY (department_id) References Departments(department_id) ON DELETE CASCADE
);
CREATE Table Admissions(
	admission_id SERIAL PRIMARY KEY,
	patient_id INT Not Null,
	doctor_id INT Null,
	room_id INT Not Null,
	FOREIGN KEY (patient_id) References Patients(patient_id) ON DELETE CASCADE,
	FOREIGN KEY (doctor_id) References Doctors(doctor_id) ON DELETE SET NULL,
	FOREIGN KEY (room_id) References Rooms(room_id) ON DELETE CASCADE,
	start_date TIMESTAMP Not Null,
	operations_number INT Not Null,
	discharge_date TIMESTAMP Null,
	status VARCHAR(20) Not Null,
	doctor_notes TEXT Null,
	price DECIMAL Not Null
);
CREATE Table Operations(
	operation_id SERIAL PRIMARY KEY,
	admission_id INT Not Null,
	room_id INT Not Null,
	patient_id INT Not Null,
	surgeon_id INT Not Null,
	FOREIGN KEY (surgeon_id) REFERENCES Doctors(doctor_id) ON DELETE CASCADE,
	FOREIGN KEY (admission_id) References Admissions(admission_id) ON DELETE CASCADE,
	FOREIGN KEY (room_id) References Rooms(room_id) ON DELETE CASCADE,
	FOREIGN KEY (patient_id) References Patients(patient_id) ON DELETE CASCADE,
	scheduled_date TIMESTAMP Not Null,
	information VARCHAR(300) Not Null,
	assistant_nurses TEXT Not Null,
	assistant_doctors TEXT Not Null,
	next_operation TIMESTAMP Null,
	finished_date TIMESTAMP Null,
	price DECIMAL Not Null,
	operation_notes TEXT Null
);
CREATE Table Payment(
	payment_id SERIAL PRIMARY KEY,
	operation_id INT Null,
	patient_id INT Not Null,
	admission_id INT Null,
	examination_id INT Null,
	FOREIGN KEY (operation_id) References Operations(operation_id) ON DELETE SET NULL,
	FOREIGN KEY (patient_id) References Patients(patient_id) ON DELETE CASCADE,
	FOREIGN KEY (admission_id) References Admissions(admission_id) ON DELETE SET NULL,
	FOREIGN KEY (examination_id) References Examinations(examination_id) ON DELETE SET NULL,
	amount DECIMAL Not Null,
	date_pay TIMESTAMP Null
);
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_email TEXT NOT NULL,
    receiver_email TEXT NOT NULL,
    title TEXT,
    content TEXT,
    timestamp TIMESTAMP DEFAULT NOW()
);

--Junction tables creation

CREATE TABLE prescriptions_medications
(
	id SERIAL PRIMARY KEY,
    prescriptions_prescription_id INT  NOT NULL,
    medications_medication_id INT NOT NULL,
	FOREIGN KEY (prescriptions_prescription_id) REFERENCES Prescriptions(prescription_id) ON DELETE CASCADE,
	FOREIGN KEY (medications_medication_id) REFERENCES Medications(medication_id) ON DELETE CASCADE,
	notes VARCHAR(500) Not Null
);

CREATE TABLE doctors_operations
(
	id SERIAL PRIMARY KEY,
    doctors_doctor_id INT NOT NULL,
    operations_operation_id INT NOT NULL,
	FOREIGN KEY (doctors_doctor_id) REFERENCES Doctors(doctor_id) ON DELETE CASCADE,
	FOREIGN KEY (operations_operation_id) REFERENCES Operations(operation_id) ON DELETE CASCADE,
	notes VARCHAR(500) Not Null
);

CREATE TABLE nurses_doctors
(
	id SERIAL PRIMARY KEY,
    nurses_nurse_id INT NOT NULL,
    doctors_doctor_id INT NOT NULL,
	FOREIGN KEY (nurses_nurse_id) REFERENCES Nurses(nurse_id) ON DELETE CASCADE,
	FOREIGN KEY (doctors_doctor_id) REFERENCES Doctors(doctor_id) ON DELETE CASCADE,
	notes VARCHAR(500) Not Null
);

--Insert statements
--All these insert statements are examples and should be treated as such, all information is randomly generated and does not
--match with real individuals credentials

INSERT Into Patients (last_name, first_name, age, gender, email, phone, address, credit)
VALUES ('Homer','Kevin John',33,'M','homerkj13@yahoo.com','0777712333','Str. Vinului 177', 1000000.00),
('Hannah','Georgia',27,'F','hannahgeo234@gmail.com','0712828737','Str. Iazului 19', 2500.00),
('Elene','Johannah',17,'F','johaelene222@gmail.com','0772837282','Str. Caraiman 87', 3000.00),
('Ivan','Andrei',19,'M','ivanoandrei@gmail.com','0782888273','Str. Morii 29',5000.00),
('Ureche','Gabriel',21,'M','urecheg123@gmail.com','0732111000','Str. Uzinei 221',5050.00),
('Lautaru','George',47,'M','lautarugeorgel2@gmail.com','0700000999','Str. Florilor 292',7400.00),
('Flowers','Andrew',23,'M','flowersandrew@gmail.com','0787234129','Str. Florariei 13',10500.00),
('Olaru','Octavian Ovidiu',39,'M','ovi3@gmail.com','0710310000','Str. Oltului 100',300.00),
('Kenneth','Josh',22,'M','kennyjosh@gmail.com','0728323123','Str. Hanlui 554',450.00),
('Timothy','Roberta',41,'F','troberta@gmail.com','0726565455','Str. Morii 18',320.00),
('Jacob','Johnathan',53,'M','jacob555@gmail.com','0744234222','Str. Delfinului 12',1.00),
('Larry','Steve',34,'M','steveyl123@gmail.com','0723843444','Str. Caraiman 20',13000.00),
('Brandon','Larry',48,'M','brandonlarry@gmail.com','0723823823','Str. Lautarului 28',500.00),
('Patrick','Alexander',37,'M','patrick8_93@gmail.com','074645633','Str. Elanului 57',1345.00),
('Harry','Robert',38,'M','harryguyrobert5@gmail.com','0723555777','Str. Garii 3',125.00);
INSERT Into Doctors (last_name, first_name, specialty, email, phone, department_id, is_Headdoctor, credit)
VALUES ('Gregory','Victor','Internist','victorgreg3@gmail.com','0788222333',3,FALSE, null),
('Jacqueline','Tanya','Osteopath','tanyajacq111@gmail.com','0711211211',6,TRUE, 0.00),
('Vincent','Hector William','Neurologist','vincent7@gmail.com','0777431000',11,TRUE, 0.00),
('Sophia','Abigail','Ophthalmologist','sophie222@gmail.com','0711111999',4,FALSE, null),
('Zachary','Karl','Psychiatrist','karlzach@gmail.com','0722334123',12,FALSE, null),
('Adams','Henry','Nephrologist','adamashen123@gmail.com','0712333912',8,FALSE, null),
('Henry','Harold','Oncologist','haroldhhh@gmail.com','0777227272',5,TRUE, 0.00),
('Frank','Logan','Urologist','loganfrank@yahoo.com','0733777333',10,TRUE, 0.00),
('Raymond','Tyler','Gastroenterologist','raytyler71@gmail.com','0739999111',1,TRUE, 0.00),
('Victoria','Julie','Orthopedic','julievictoria@yahoo.com','0711100239',2,TRUE, 0.00),
('Kathie','Debra','Pulmonologist','kathdebra@gmail.com','0712333129',14,FALSE, null),
('Samuel','Benjamin','Cardiologist','benjisam6@gmail.com','0723888888',7,FALSE, null),
('Aaron','Jack','Radiologist','jackthea@gmail.com','0722228989',9,TRUE, 0.00),
('Kyle','Joey','Endocrinologist','joeykyle@gmail.com','07123890475',13,FALSE, null),
('Winters','Ethan','Dermatologists','ethanwin333@gmail.com','0700000101',15,TRUE, 0.00),
('Jeffrey','James Mark','Internist','jeffjames@outlook.com','0777777878',3,TRUE,0.00),
('Mark','Tyler','Internist','tylerm2@gmail.com','0789898989',3,FALSE, null),
('Terry','Willie','Osteopath','willter@gmail.com','0726333726',6,FALSE, null),
('Ben','Benny','Osteopath','bennys@gmail.com','0762323129',6,FALSE, null),
('Eric','Larry Johnathan','Neurologist','ericsen@gmail.com','0737722120',11,FALSE, null),
('Luther','Scott Jacob','Neurologist','lutherjascot@gmail.com','0799222667',11,FALSE, null),
('Stephen','Jacob','Ophthalmologist','stephenjac@gmail.com','0799299999',4,FALSE, null),
('Kevin','Octavian','Ophthalmologist','octak@gmail.com','0723882888',4,TRUE,0.00),
('Noah','Greg','Psychiatrist','gregnoah@yahoo.com','0777777770',12,TRUE,0.00),
('Sam','Samuel','Psychiatrist','samsam@gmail.com','0728222222',12,FALSE, null),
('Nistor','Jeremy','Nephrologist','jern@gmail.com','0777777710',8,TRUE,0.00),
('Brandon','Chris','Nephrologist','chrisbrandon@gmail.com','0723222000',8,FALSE, null),
('Austin','Robert','Oncologist','austinrob@gmail.com','0700002111',5,FALSE, null),
('Paul','Octavian','Oncologist','paulocta@gmail.com','0789898222',5,FALSE, null),
('Dylan','Henry','Urologist','henryd@gmail.com','0722222220',10,FALSE, null),
('Gerald','Harold','Urologist','gerharold@gmail.com','0700000001',10,FALSE, null),
('Victor','Robert','Gastroenterologist','victor@gmail.com','0788888001',1,FALSE, null),
('Jesse','Harold','Gastroenterologist','jesharold@gmail.com','0799191212',1,FALSE, null),
('Jackie','John','Orthopedic','johnjack@gmail.com','0719999991',2,FALSE, null),
('Lawrence','Kevin','Orthopedic','kevlaw@gmail.com','0722888110',2,FALSE, null),
('Willie','Mason Lucas','Pulmonologist','masonwill@gmail.com','0700101111',14,TRUE,0.00),
('Larry','Arthur','Pulmonologist','larryart@gmail.com','0788888880',14,FALSE, null),
('Morgan','Arthur','Cardiologist','arthurmorgan@yahoo.com','0700000110',7,TRUE, 0.00),
('Lawrence','Jack','Cardiologist','jacklaw@gmail.com','0788222383',7,FALSE, null),
('Russell','John','Radiologist','johnruss@gmail.com','0722811111',9,FALSE, null),
('Cooper','Will','Radiologist','cooperw@gmail.com','0711111111',9,FALSE, null),
('Wayne','Walter Will','Endocrinologist','walterwww@gmail.com','0733777000',13,TRUE, 0.00),
('Ericsen','Arthur','Endocrinologist','erica@gmail.com','0711111000',13,FALSE, null),
('Carl','Jordan','Dermatologists','jordank@gmail.com','0700999222',15,FALSE, null),
('Kevin','Dylan','Dermatologists','kevdylan@gmail.com','0799888005',15,FALSE, null);
INSERT Into Medications (name_medication, description, price, department_id, inventory)
VALUES ('Naproxen','Anti-inflammatory drug for arthritis, muscle aches',117.00,6,3),
('Meloxicam','Anti-inflammatory drug, reduces swelling and pain in joints',300.00,6,11),
('Cyclobenzaprine','Muscle relaxant, for musculoskeletal pain and spasms',210.00,6,8),
('Gabapentin','Used for nerve pain, epilepsy, seizures',550.70,11,3),
('Levetiracetam','Unique antiepileptic drug, which controls seizures',717.50,11,1),
('Amitriptyline','Helps with chronic nerve pain and migraines',330.00,11,4),
('Latanoprost','Medication for glaucoma',40.50,4,11),
('Timolol','Treats eye pressure, glaucoma',70.00,4,7),
('Olopatadine','Used for allergic conjunctivitis',89.00,4,9),
('Sertraline','Medication for depression, anxiety – SSRI',130.00,12,9),
('Olanzapine','Treats schizophrenia, bipolar disorder',470.00,12,3),
('Diazepam','Used for anxiety, panic attacks',350,12,7),
('Furosemide','Diuretic for fluid retention',70.00,8,11),
('Erythropoietin','Treats anemia from kidney failure',110.00,8,6),
('Sodium Bicarbonate','Corrects metabolic acidosis',50.00,8,7),
('Cisplatin','Medication for chemotherapy',140.00,5,20),
('Paclitaxel','Used for breast, ovarian, lung cancers',250.00,5,11),
('Bevacizumab','Medication for anti-angiogenesis therapy',300.00,5,5),
('Tamsulosin','For benign prostatic hyperplasia',55.50,10,7),
('Oxybutynin','For overactive bladder',135.00,10,5),
('Finasteride','For prostate enlargement',700.00,10,1),
('Omeprazole','Used for acid reflux, ulcers – PPI',340.00,1,3),
('Mesalamine','Treats inflammatory bowel disease',120.00,1,10),
('Rifaximin','For for IBS, liver-related encephalopathy',50.00,1,4),
('Ibuprofen','Medication used for pain relief. Its anti-inflammatory',30.99,2,30),
('Alendronate','For osteoporosis',77.00,2,9),
('Celecoxib','Used for arthritis pain',100.00,2,5),
('Salbutamol','For asthma inhaler',40.50,14,9),
('Fluticasone','Medication for inhaled corticosteroid',30.99,14,7),
('Montelukast','Used for asthma and allergies',390.00,14,5),
('Amlodipine','Medication for high blood pressure',140.00,7,13),
('Clopidogrel','It prevents blood clots',111.11,7,9),
('Atorvastatin','For cholesterol control',450.00,7,3),
('Iohexol','It is a iodine-based contrast for CT scans',133.33,9,4),
('Gadobutrol','Gadolinium-based contrast for MRI',421.00,9,5),
('Barium Sulfate','Contrast agent for GI X-rays',500.00,9,7),
('Metformin','First-line for Type 2 diabetes',70.00,13,10),
('Levothyroxine','Hypothyroidism treatment',777.00,13,1),
('Hydrocortisone','Adrenal insufficiency, cortisol replacement',100.00,13,11),
('Isotretinoin','For severe acne treatment',300.00,15,5),
('Clobetasol Propionate','Strong corticosteroid for eczema, psoriasis',70.00,15,10),
('Mupirocin','Antibiotic cream for skin infections',20.50,15,30),
('Paracetamol','Used for pain and fever relief',20.00,3,7),
('Amoxicillin','Broad-spectrum antibiotic',35.70,3,11),
('Lisinopril','ACE inhibitor for blood pressure',55.00,3,9);
INSERT Into Departments (name, rooms_number, description)
VALUES ('Gastroenterology',6,'Focuses on the digestive system and its disorders.'),
('Orthopedy',6,'Deals with conditions involving the muscle/skeleton system.'),
('Internal Medicine',6,'Focuses on the diagnosis, treatment, and prevention of adult diseases, particularly chronic conditions affecting internal organs.'),
('Ophthalmology',6,'Specializes in the treatment of eye conditions.'),
('Oncology',6,'Focuses on the diagnosis and treatment of cancer.'),
('Osteopathy',6,'Emphasizes physical manipulation of the body to treat various conditions.'),
('Cardiology',6,'Deals with heart diseases and circulatory system issues.'),
('Nephrology',6,'Focuses on kidney-related diseases and treatment.'),
('Radiology',6,'Specializes in diagnostic imaging like X-rays, CT scans, and MRIs.'),
('Urology',6,'Deals with the urinary tract and male reproductive organs.'),
('Neurology',6,'Focuses on disorders of the nervous system.'),
('Psychiatry',6,'Focuses on mental health disorders and treatment.'),
('Endocrinology',6,'Deals with hormone disorders and diseases.'),
('Pulmonology',6,'Specializes in lung diseases and respiratory issues.'),
('Dermatology',6,'Focuses on skin conditions and diseases.');
INSERT Into Rooms (number_room, department_id, type_room, capacity)
VALUES (1, 1, 'Examination Room', 5),
(2, 1, 'Admitted Patients Room', 15),
(3, 1, 'Operation Room', 7),
(4, 1, 'Staff Room', 6),
(5, 1, 'Waiting Room', 10),
(6, 1, 'Head Doctors Office', 1),
(7, 2, 'Examination Room', 5),
(8, 2, 'Admitted Patient Room', 15),
(9, 2, 'Operation Room', 7),
(10, 2, 'Staff Room', 6),
(11, 2, 'Waiting Room', 10),
(12, 2, 'Head Doctors Office', 1),
(13, 3, 'Examination Room', 5),
(14, 3, 'Admitted Patient Room', 15),
(15, 3, 'Operation Room', 7),
(16, 3, 'Staff Room', 6),
(17, 3, 'Waiting Room', 10),
(18, 3, 'Head Doctors Office', 1),
(19, 4, 'Examination Room', 5),
(20, 4, 'Admitted Patient Room', 15),
(21, 4, 'Operation Room', 7),
(22, 4, 'Staff Room', 6),
(23, 4, 'Waiting Room', 10),
(24, 4, 'Head Doctors Office', 1),
(25, 5, 'Examination Room', 5),
(26, 5, 'Admitted Patient Room', 15),
(27, 5, 'Operation Room', 7),
(28, 5, 'Staff Room', 6),
(29, 5, 'Waiting Room', 10),
(30, 5, 'Head Doctors Office', 1),
(31, 6, 'Examination Room', 5),
(32, 6, 'Admitted Patient Room', 15),
(33, 6, 'Operation Room', 7),
(34, 6, 'Staff Room', 6),
(35, 6, 'Waiting Room', 10),
(36, 6, 'Head Doctors Office', 1),
(37, 7, 'Examination Room', 5),
(38, 7, 'Admitted Patient Room', 15),
(39, 7, 'Operation Room', 7),
(40, 7, 'Staff Room', 6),
(41, 7, 'Waiting Room', 10),
(42, 7, 'Head Doctors Office', 1),
(43, 8, 'Examination Room', 5),
(44, 8, 'Admitted Patient Room', 15),
(45, 8, 'Operation Room', 7),
(46, 8, 'Staff Room', 6),
(47, 8, 'Waiting Room', 10),
(48, 8, 'Head Doctors Office', 1),
(49, 9, 'Examination Room', 5),
(50, 9, 'Admitted Patient Room', 15),
(51, 9, 'Operation Room', 7),
(52, 9, 'Staff Room', 6),
(53, 9, 'Waiting Room', 10),
(54, 9, 'Head Doctors Office', 1),
(55, 10, 'Examination Room', 5),
(56, 10, 'Admitted Patient Room', 15),
(57, 10, 'Operation Room', 7),
(58, 10, 'Staff Room', 6),
(59, 10, 'Waiting Room', 10),
(60, 10, 'Head Doctors Office', 1),
(61, 11, 'Examination Room', 5),
(62, 11, 'Admitted Patient Room', 15),
(63, 11, 'Operation Room', 7),
(64, 11, 'Staff Room', 6),
(65, 11, 'Waiting Room', 10),
(66, 11, 'Head Doctors Office', 1),
(67, 12, 'Examination Room', 5),
(68, 12, 'Admitted Patient Room', 15),
(69, 12, 'Operation Room', 7),
(70, 12, 'Staff Room', 6),
(71, 12, 'Waiting Room', 10),
(72, 12, 'Head Doctors Office', 1),
(73, 13, 'Examination Room', 5),
(74, 13, 'Admitted Patient Room', 15),
(75, 13, 'Operation Room', 7),
(76, 13, 'Staff Room', 6),
(77, 13, 'Waiting Room', 10),
(78, 13, 'Head Doctors Office', 1),
(79, 14, 'Examination Room', 5),
(80, 14, 'Admitted Patient Room', 15),
(81, 14, 'Operation Room', 7),
(82, 14, 'Staff Room', 6),
(83, 14, 'Waiting Room', 10),
(84, 14, 'Head Doctors Office', 1),
(85, 15, 'Examination Room', 5),
(86, 15, 'Admitted Patient Room', 15),
(87, 15, 'Operation Room', 7),
(88, 15, 'Staff Room', 6),
(89, 15, 'Waiting Room', 10),
(90, 15, 'Head Doctors Office', 1);
INSERT INTO Nurses (last_name, first_name, phone, department_id)
VALUES ('Carter','Emily','0711111222',1),
('Brooks','Olivia','0722111222',1),
('Simmons','Mia','0711211111',1),
('Murphy','Emma','0711888222',2),
('Hayes','Sophia','0788282828',2),
('Bailey','Lily','0711828282',2),
('Wilson','James','0712727272',3),
('Mitchell','Noah','0711921929',3),
('Rogers','Mason','0711929929',3),
('Reed','Lucas','0711989383',4),
('Sullivan','Ryan','0711889332',4),
('Bryant','Chloe','0711999883',4),
('Anderson','Caleb','0711888223',5),
('Parker','Grace','0722888333',5),
('Bennett','Aiden','0711829333',5),
('Russell','Ethan','0766399440',6),
('Roberts','Ava','0711009334',6),
('Cooper','Isabella','0711900002',6),
('Hayes','Logan','0711920334',7),
('Bryant','Emily','0713123123',7),
('Carter','Sophia','0788999999',7),
('Brooks','Ryan','0711111123',8),
('Hayes','Lily','0712111111',8),
('Bailey','Caleb','0711123141',8),
('Reed','Olivia','0722993933',9),
('Wilson','Ethan','0733883883',9),
('Simmons','Grace','0733444444',9),
('Murphy','Mason','0733444447',10),
('Rogers','Lucas','0733444448',10),
('Sullivan','Chloe','0733444449',10),
('Reed','Aiden','0733999223',11),
('Bailey','Noah','0733812938',11),
('Russell','Emma','0734982344',11),
('Cooper','Mia','0733812934',12),
('Bryant','Isabella','0755923000',12),
('Hayes','James','0755000011',12),
('Bennett','Ava','0755999555',13),
('Roberts','Ryan','0755832495',13),
('Anderson','Sophia','0755738237',13),
('Parker','Lily','0755834892',14),
('Brooks','Caleb','0722838955',14),
('Simmons','Olivia','0788238924',14),
('Murphy','Ethan','0788213835',15),
('Rogers','Grace','0755828883',15),
('Sullivan','Mason','0755823823',15);
INSERT Into Appointments(patient_id, doctor_id, room_id, schedule_date, status)
VALUES (3,5,67,'2025-07-11 15:00:00','waiting'),
(5,5,67,'2025-04-30 17:40:00','scheduled'),
(5,11,79,'2025-05-01 11:30:00','scheduled'),
(3,11,79,'2025-04-27 20:23:00','scheduled'),
(12,18,31,'2025-05-01 11:00:00','scheduled'),
(8,43,73,'2025-02-17 15:00:00','examined'),
(2,43,73,'2024-12-07 15:00:00','examined'),
(2,43,73,'2025-01-22 15:00:00','examined'),
(2,43,73,'2025-02-17 13:00:00','examined'),
(13,27,43,'2024-03-03 18:40:00','examined'),
(13,43,73,'2025-03-04 19:55:00','examined'),
(10,27,43,'2025-09-24 11:23:00','examined'),
(9,5,67,'2024-12-07 10:35:00','examined'),
(10,1,13,'2025-01-22 09:00:00','examined'),
(9,1,13,'2025-02-17 20:00:00','examined'),
(3,1,13,'2024-12-07 13:10:00','examined'),
(3,1,13,'2025-01-22 12:17:00','examined'),
(4,5,67,'2025-03-19 15:00:00','canceled'),
(3,43,73,'2024-07-08 15:00:00','canceled'),
(12,1,13,'2025-04-10 15:00:00','canceled'),
(4,9,1,'2025-04-05 14:35:00', 'examined'),
(5,9,1,'2025-04-10 14:00:00', 'examined'),
(6,9,1,'2025-04-11 13:00:00', 'examined'),
(15,39,37,'2025-02-17 10:00:00','examined'),
(15,45,85,'2025-03-13 10:00:00','examined'),
(15,43,73,'2025-01-01 10:00:00','examined'),
(13,39,37,'2024-11-17 10:00:00','examined'),
(12,45,85,'2024-09-07 10:00:00','examined'),
(11,45,85,'2025-04-04 10:00:00','examined'),
(10,39,37,'2025-04-05 10:00:00','examined'),
(9,20,61,'2025-02-14 10:00:00','examined');
INSERT Into Examinations (appointment_id, examination_date, information, assistant_nurses, diagnosis, result, reexamination_date, price, level, examination_notes)
VALUES
(6,'2025-02-17 15:30:00','Olaru Octavian Ovidiu - thirsty often, nausea during the day','Ava, Ryan & Sophia','Diabetes','admission',null,30.00,'moderate','Examination 1: Patient seems to suffer from diabetes after drinking water and still being thirsty and having nausea which are diabetes symptoms.'),
(7,'2024-12-07 15:00:00','Hannah Georgia - fatigue','Ava & Ryan','Unknown','reexamination','2025-01-22 15:00:00',30.00,'none','Examination 1: Patient seems to suffer from adrenal insufficiency after being tested, having fatigue, should reexamine later.'),
(8,'2025-01-22 15:00:00','Hannah Georgia - fatigue','Ava & Ryan','Unknown','reexamination','2025-02-17 13:00:00',70.00,'none','Examination 2: Patient seems to suffer from adrenal insufficiency after being tested, having fatigue, weakness and loss of appetite, should reexamine later.'),
(9,'2025-02-17 13:00:00','Hannah Georgia - fatigue,weakness,loss of appetite','Ava & Ryan','Adrenal insufficiency','admission',null,110.00,'moderate','Examination 3: Patient suffers from adrenal insufficiency after being tested, having fatigue, weakness and loss of appetite, she is admissed to the endocrynology department.'),
(10,'2024-03-03 18:40:00','Brandon Larry - frequent urine,fever,chills','Ryan, Lily & Caleb','Kidney Infection','prescription',null,45.00,'mild','Examination 1: Patient seems to have kidney infections after the urine test, loss of appetite and fever tests the examination concluded in a prescription he should follow for a month.'),
(11,'2025-03-04 19:55:00','Brandon Larry - fatigue,weakness','Ava & Ryan','Adrenal insufficiency','admission',null,30.00,'moderate','Examination 1: Patient suffers from adrenal insufficiency after being tested, having fatigue, weakness and loss of appetite, he is admissed to the endocrynology department.'),
(12,'2025-09-24 11:23:00', 'Timothy Roberta - High blood pressure', 'Ryan & Lily','Polycystic Kidney Disease','admission',null,110.00,'critical','Examination 1: Patient has high blood pressure, suspected to have kidney stones and headaches, should be admissed immediately!'),
(13,'2024-12-07 10:35:00', 'Kenneth Josh - forgets quick, sees things others dont, cant speak properly', 'Isabella, Mia & James','Schizophrenia','admission',null,1000.00,'severe','Examination 1: Patient seems to have delusions, uncontrolled impulses, speech impediment and memory loss, he should be admissed immediately having symptoms of schizophrenia!'),
(14,'2025-01-22 09:00:00', 'Timothy Roberta - headaches', 'James, Noah & Mason','Migraines','admission',null,50.00,'moderate','Examination 1: Patient is examined inside Internal Medicine, presenting symptoms of an advanced migraine, she is admissed inside the Osteopathy department for treatment.'),
(15, '2025-02-17 20:00:00', 'Kenneth Josh - painful swallowing', 'James & Noah','Strep Throat','prescription',null,35.00,'mild','Examination 1: Patient has issues swallowing, presenting throat pain and swollen lympth nodes, likely strep throat.'),
(16, '2024-12-07 13:10:00' ,'Elene Johannah - bone pain', 'James', 'Unknown','reexamination', '2025-01-22 12:17:00',30.00,'none','Examination 1: Patient informs that she is feeling pain in the bone, specifically the left leg, after investigations we found a lump, should be reexamined in a month.'),
(17, '2025-01-22 12:17:00', 'Elene Johannah - bone pain', 'James, Noah & Mason', 'Sarcomas', 'admission', null, 500.00, 'terminal','Examination 2: Patient suffered extreme weight loss, the lump has grown considerably large and still suffers from bone pain, this is unfortunately sarcomas and patient needs intensive care immediately!'),
(21,'2025-04-05 14:35:00', 'Ivan - burning in chest', 'Emily, Olivia & Mia','Gastroesophageal reflux disease','admission',null,200.00,'moderate','Examination 1: Patient has been examined and we found out hes suffering from heartburn, symptom of gastroesophageal reflux disease.'),
(22,'2025-04-10 14:00:00', 'Ureche Gabriel - diarrhea','Emily & Olivia','Coeliac Disease','admission',null,320.00,'moderate','Examination 1: Patient is suffering from diarrhea and and abdominal pain, this leads to further analysis which results in coeliac disease.'),
(23,'2025-04-11 13:00:00', 'Lautaru George - weight loss,abdominal pain','Emily & Mia','Inflammatory Bowel Disease','admission',null,510.00,'moderate','Examination 1: Patient is suffering from weight loss, internal bleeding and abdominal pain, symptoms clearly of Crohns disease or inflammatory bowel disease.'),
(24,'2025-02-17 10:00:00','Harry Robert - dizziness','Logan,Emily & Sophia','Abnormal heart rhythms','prescription',null,50.00,'mild','Examination 1: Patient suffers from dizziness and heart palpitations, he needs prescription.'),
(25,'2025-03-13 10:00:00','Harry Robert - red and white spots','Ethan,Grace & Mason','Acne','prescription',null,35.00,'mild','Examination 1: Patient presents himself with red and white spots on his facial skin, meaning acne.'),
(26,'2025-01-01 10:00:00','Harry Robert - rapid heartbeat, weakness','Ava,Ryan & Sophia','Hyperthyroidism','prescription',null,60.00,'mild','Examination 1: Patient seems to suffer from weight loss since hes feeling weaker and also has a rapid heartbeat, this is hyperthyroidism.'),
(27,'2024-11-17 10:00:00','Brandon Larry - heavy breath,fatigue','Logan & Emily','Valvular heart disease','prescription',null,75.00,'mild','Examination 1: Patient suffers from shortness of breath and fatigue, meaning valvular heart disease.'),
(28,'2024-09-07 10:00:00','Larry Steve - itching','Ethan & Grace','Candidiasis','prescription',null,80.00,'mild','Examination 1: Patient has itching and a red patch on the skin, meaning candidiasis.'),
(29,'2025-04-04 10:00:00','Jacob - red skin','Ethan','Impetigo','prescription',null,50.00,'mild','Examination 1: Patient has impetigo symptoms with reddish sores patches. We recommended a prescription he should follow.'),
(30,'2025-04-05 10:00:00','Roberta - heavey breath,fainting','Logan','Rheumatic heart disease','prescription',null,78.00,'mild','Examination 1: Patient has shortness of breath and fainting symptoms.'),
(31,'2025-02-14 10:00:00','Josh - numbness and headaches','Aiden,Noah & Emma','Vasculitis Syndromes','prescription',null,45.00,'mild','Examination 1: Patient has numbness in the limbs, headaches and dizziness, this is vasculitis syndrome which affects the central and peripherous nervous system.');
INSERT Into Admissions (patient_id, doctor_id, room_id, start_date, operations_number, discharge_date, status, doctor_notes, price)
VALUES (8, 43, 74, '2025-02-18 08:30:00',1,'2025-02-25 22:00:00','discharged','Olaru went into investigation processes with his diabetes and was being given insulin injections frequently.',1000.00),
(2, 14, 74, '2025-02-18 13:00:00',2,'2025-02-25 22:00:00','discharged','Hannah went into investigation with the adrenal insufficiency being given hormone medicines for adrenal glands.',300.00),
(13, 14, 74, '2025-03-05 09:00:00',0,'2025-03-10 22:00:00','discharged','Larry went into investigation with the adrenal insufficiency being given hormone medicines for adrenal glands.',300.00),
(10, 27, 44, '2025-09-25 11:23:00', 3, '2025-10-10 22:00:00','discharged','Roberta went into investigation with her Polycystic Kidney Disease being given pills to reduce the cysts growth.',170.00),
(9, 25, 68, '2024-12-07 19:35:00', 1, '2024-12-30 22:35:00', 'discharged','Kenneth went into investigation with his schizophrenia and was being given all sorts of antipsychotic  medications.',570.00),
(10, 18, 32, '2025-01-23 09:00:00', 2, '2025-02-01 22:00:00', 'discharged','Roberta went into investigation with her headaches issues which developed in advanced migraines being given anti-inflammatory medicine.',300.00),
(3, 29, 26, '2025-01-22 15:17:00', 5, '2025-02-28 22:00:00', 'discharged','Elene went into investigation with the sarcomas and did radiation therapy and took pills to treat the cancers.',1500.00),
(4,9,2,'2025-04-05 14:35:00',0,null,'active','Ivan went into investigation with the gastroesophageal reflux disease and is being given special medicines for treatment.',700.00),
(5,9,2,'2025-04-10 14:00:00',0,null,'active','Ureche Gabriel went into investigation with the coeliac disease treated with gluten-free foods and medications.',500.00),
(6,9,2,'2025-04-11 13:00:00',1,null,'active','Lautaru George is being investigated into medical admission with the Crohns disease and is being given medications to reduce inflammations.',1100.00);
INSERT Into Operations (admission_id, room_id, patient_id, surgeon_id, scheduled_date, information, assistant_nurses, assistant_doctors, next_operation, finished_date, price, operation_notes)
VALUES
(1,75,8,42,'2025-02-22 12:30:00','Olaru went into surgery for the gallbladder removal.','Ava, Ryan & Sophia','Joey & Arthur',null,'2025-02-22 13:30:00',300.00,'Olaru did the cholecystectomy operation in which his gallbladder was succesfully removed, curing him of diabetes!'),
(2,75,2,42,'2025-02-20 12:00:00','Hannah went into surgery for surface hormone injection.','Ava & Ryan', 'Joey','2025-02-21 12:00:00','2025-02-20 12:50:00',100.00,'Hannah was given extremely high dosages of hormones for her insufficiency in an operation. Theres one more operation that needs to be done.'),
(2,75,2,42,'2025-02-21 12:00:00','Hannah went into surgery for surface hormone injection.','Ava & Ryan', 'Joey',null,'2025-02-21 12:50:00',100.00,'Hannah was given extremely high dosages of hormones for her insufficiency in an operation, She is cured!.'),
(4,45,10,26,'2025-09-26 11:23:00','Roberta went into the first surgery for kidney cyst removal.','Ryan,Lily & Caleb','Henry & Chris','2025-09-27 11:23:00','2025-09-26 13:23:00', 300.00,'Roberta has parts of her kidney removed since they contained cysts. Scheduled for the second operation.'),
(4,45,10,26,'2025-09-27 11:23:00','Roberta went into the second surgery for kidney cyst removal.','Ryan,Lily & Caleb','Henry & Chris','2025-10-03 11:23:00','2025-09-27 13:23:00', 300.00,'Roberta has parts of her kidney removed since they contained cysts. Scheduled for the third operation.'),
(4,45,10,26,'2025-10-03 11:23:00','Roberta went into the third surgery for kidney cyst removal.','Ryan,Lily & Caleb','Henry & Chris',null,'2025-10-03 13:23:00', 300.00,'Roberta has parts of her kidney removed since they contained cysts. She is finally cured!.'),
(5,69,9,24,'2024-12-27 19:35:00','Kenneth went into the deep brain stimulation surgery, where electrodes will be placed inside his brain.','Isabella,Mia & James','Samuel & Karl',null,'2024-12-28 03:00:00',3500.00,'Kenneth has been succesfully treated from his schizofrenia using the deep brain stimulation surgery, where electrodes were placed in side his brain for manual configurations to alter his behaviours.'),
(6,33,10,2,'2025-01-25 11:00:00','Roberta went into her first osteopathy massage therapy for her migraines.','Isabella,Ava & Ethan','Willie & Benny','2025-01-30 11:00:00','2025-01-25 12:00:00',370.00,'Roberta has been given the craniosacral massage which improved her overall moods and movement functions. Next therapy is scheduled.'),
(6,33,10,2,'2025-01-30 11:00:00','Roberta went into her first osteopathy massage therapy for her migraines.','Isabella,Ava & Ethan','Willie & Benny',null,'2025-01-30 13:00:00',370.00,'Roberta has been given the craniosacral massage which improved her overall moods and movement functions. She is relaxed and cured from the migraines.'),
(7,27,3,7,'2025-01-30 17:10:00','Elene went into an invasive operation for tissue removal.','Aiden,Grace & Caleb','Robert & Octavian','2025-02-03 17:00:00','2025-01-30 19:00:00',750.00,'Elene went into the invasive tissue removal operation, which removes parts of cancerous tissues from her leg, final goal being removing the whole sarcoma. Second surgery scheduled.'),
(7,27,3,7,'2025-02-03 17:00:00','Elene went into the second operation for tissue removal.','Aiden,Grace & Caleb','Robert & Octavian','2025-02-07 17:00:00','2025-02-03 19:00:00',750.00,'Elene went into the invasive tissue removal operation, which removes parts of cancerous tissues from her leg, final goal being removing the whole sarcoma. Third surgery scheduled.'),
(7,27,3,7,'2025-02-07 17:00:00','Elene went into the third operation for tissue removal.','Aiden,Grace & Caleb','Robert & Octavian','2025-02-17 17:00:00','2025-02-07 19:00:00',750.00,'Elene went into the invasive tissue removal operation, which removes parts of cancerous tissues from her leg, final goal being removing the whole sarcoma. Fourth surgery scheduled.'),
(7,27,3,7,'2025-02-17 17:00:00','Elene went into the fourth operation for tissue removal.','Aiden,Grace & Caleb','Robert & Octavian','2025-02-23 17:00:00','2025-02-17 19:00:00',750.00,'Elene went into the invasive tissue removal operation, which removes parts of cancerous tissues from her leg, final goal being removing the whole sarcoma. Fifth surgery scheduled.'),
(7,27,3,7,'2025-02-23 17:00:00','Elene went into the fifth and final operation for tissue removal.','Aiden,Grace & Caleb','Robert & Octavian',null,'2025-02-23 19:00:00',750.00,'Elene went into the invasive tissue removal operation, which removes parts of cancerous tissues from her leg, final goal achieved, the sarcomas have been completely removed from her organism. She is finally and successfully cured!'),
(10,2,6,9,'2025-04-25 12:30:00','Lautaru went into the surgical operation of repairing the affected gastrointestinal tract.','Emily,Olivia & Mia','Robert & Harold',null,'2025-04-25 15:30:00',730.30,'Lautaru went into the gastrointestinal tract repairing operation, which will make him recover in a few weeks from the disease, being completely cured.');
INSERT Into Prescriptions (examination_id, number_medications, notes, prescription_date, total_price)
VALUES (5, 2, 'Prescription: Furosemide 3 pill a day & Erythropoietin 1 pill', '2024-03-03 18:55:00',180.00),
(10, 3, 'Prescription: Amoxicillin 3 antibiotics a day, Lisinopril 3 tbsp a day & Paracetamol 1 pill a day','2025-02-17 20:30:00',110.70),
(16,1,'Prescription: Amlodipine 3 pills a day','2025-02-17 10:30:00',140.00),
(17,1,'Prescription: Isotretinoin 3 pills a day for acne treatment','2025-03-13 10:00:00',300.00),
(18,1,'Prescription: Levothyroxine 5 pills a day for hormone tratment','2025-01-01 10:00:00',777.00),
(19,3,'Prescription: Atorvastatin 3 pills a day, Clopidogrel 2 tbsp & Amlodipine 2 pills a day','2024-11-17 10:00:00',701.11),
(20,2,'Prescription: Clobetasol Propionate 2 pills a day & Mupirocin 1 tbsp a day','2024-09-07 10:00:00',90.50),
(21,1,'Prescription: Mupirocin 3 tbsp a day','2025-04-04 10:00:00',20.50),
(22,1,'Prescription: Amlodipine 2 pills a day','2025-04-05 10:00:00',140.00),
(23,3,'Prescription: Gabapentin 3 tbsp a day & Amitriptyline 3 pills a day','2025-02-14 10:00:00',1598.20);
INSERT Into Payment (operation_id, patient_id, admission_id, examination_id,amount,date_pay)
VALUES (2,2,2,2,430.00,'2025-02-25 21:00:00'),
(3,2,null,3,170.00,'2025-02-25 21:00:00'),
(null,2,null,4,110.00,'2025-02-25 21:00:00'),
(7,9,5,8,5070.00,'2024-12-30 20:35:00'),
(null,9,null,10,145.70,'2024-12-30 21:35:00'),
(4,10,4,7,580.00,'2025-10-10 21:00:00'),
(5,10,6,9,650.00,'2025-10-10 21:10:00'),
(6,10,null,22,518.00,'2025-10-10 21:20:00'),
(8,10,null,null,370.00,'2025-10-10 21:30:00'),
(9,10,null,null,370.00,'2025-10-10 21:45:00');
INSERT Into nurses_doctors (nurses_nurse_id, doctors_doctor_id, notes)
VALUES (37,43,'Ava prepped the ustensils needed for Ovidius first examination.'),
(38,43,'Ryan injected a chemical into Ovidius circulatory system to test hormone reaction.'),
(39,43,'Sophia analysed Ovidius vital signs and recorded his hormone reaction to the injected chemical, Ovidiu having diabetes.'),
(37, 43, 'Ava monitored Hannahs blood pressure and a injected a special chemical for adrenal testing during first exam.'),
(38, 43, 'Ryan analysed vital signs and documented hormone response for adrenal function check.'),
(37, 43, 'Ava assisted with the second reexamination, recorded patient fatigue and loss of appetite.'),
(38, 43, 'Ryan took blood samples for hormonal panel and verified hormonal levels.'),
(37, 43, 'Ava helped prepare Hannah for endocrinology admission, managing fluids and patient comfort.'),
(38, 43, 'Ryan performed final vital check and ensured hormone analysis data was sent to lab.'),
(34, 5, 'Mia prepared the cognitive test for Kenneth and registered his verbal reactions.'),
(35, 5, 'Isabella registered his cognitive responses along with his behaviour patterns.'),
(36, 5, 'James managed the reflex test and analysed the psychiatric examination.'),
(43, 45, 'Ethan applied antiseptic and prepared acne lesion sites for examination.'),
(44, 45, 'Grace documented the red spots on the lesion sites while assisting Dr. Dylan during the skin inspection.'),
(45, 45, 'Mason conducted a patch test for allergic reactions and went into council with the nurses for the prescription.'),
(31, 20, 'Aiden recorded neurological responses and managed the coordination tests.'),
(32, 20, 'Noah administered the blood pressure and vasculitis panel tests, reporting abnormalities.'),
(33, 20, 'Emma monitored patient vitals during the procedure.'); 
INSERT Into doctors_operations (doctors_doctor_id, operations_operation_id, notes)
VALUES (42,1,'Headdoctor Walter Will performed the laparoscopic cholecystectomy and supervised the surgical team.'),
(14,1,'Joey assisted by retracting the tissues and monitoring internal bleeding.'),
(43,1, 'Arthur provided instrument support and maintained visibility in the operating theatre.'),
(24, 7,'Headdoctor Greg performed the DBS implantation and supervised electrode calibration.'),
(5, 7,'Karl assisted during the electrode placement'),
(25, 7,'Samuel monitored the patients neural responses.'),
(9,15,'Headdoctor Tyler led the gastrointestinal tract reconstruction and managed the surgical plan'),
(32, 15, 'Robert supported with suturing.'),
(33,15,'Harold monitored bleeding control and maintained the sterile field.'),
(7, 14, 'Dr. Henry Harold, lead oncologist, successfully directed the final sarcoma removal surgery on Elenes left leg. Oversaw all critical phases of the operation.'),
(28, 14, 'Robert assisted in tumor removal and ensured no residual infected tissue remained.'),
(29, 14, 'Octavian handled post-removal tissue management, helping complete the final removal phase efficiently.'),
(14,14,'Grace, nurse, prepped surgical instruments, monitored Elene’s vital signs, and assisted in maintaining a sterile environment throughout the sarcoma removal operation.');
INSERT Into prescriptions_medications (prescriptions_prescription_id, medications_medication_id, notes)
VALUES (1,13,'Furosemide prescribed in prescription with ID 1 for patient Brandon Larry with Kidney Infection.'),
(1,14,'Erythropoietin administered in prescription with ID 1 for patient Brandon Larry with Kidney Infection.'),
(2,43,'Paracetamol prescribed in prescription with ID 2 for patient Kenneth Josh with Strep Throat.'),
(2,44,'Amoxicillin prescribed in prescription with ID 2 for patient Kenneth Josh with Strep Throat.'),
(2,45,'Lisinopril prescribed in prescription with ID 2 for patient Kenneth Josh with Strep Throat.'),
(3,31,'Amlodipine prescribed in prescription with ID 3 for patient Harry Robert with Abnormal heart rhythms.'),
(4,40,'Isotretinoin prescribed in prescription with ID 4 for patient Harry Robert with Acne.'),
(5,38,'Levothyroxine prescribed in prescription with ID 5 for patient Harry Robert with Hyperthyroidism.'),
(6,31,'Amlodipine prescribed in prescription with ID 6 for patient Brandon Larry with Valvular heart disease.'),
(6,32,'Clopidogrel prescribed in prescription with ID 6 for patient Brandon Larry with Valvular heart disease.');

--Joins

--Information about patients and their doctors through appointments, because doctors do not have a foreign key with patients
SELECT Patients.patient_id, Patients.last_name, Patients.first_name,
	Doctors.last_name, Doctors.first_name, Doctors.specialty, Doctors.is_Headdoctor,
	Examinations.information
From Patients
INNER JOIN Appointments on Appointments.patient_id=Patients.patient_id
INNER JOIN Doctors on Doctors.doctor_id=Appointments.doctor_id
INNER JOIN Examinations on Appointments.appointment_id=Examinations.appointment_id
ORDER By Examinations.examination_date;

--Lista de medicamente impreuna cu prescriptii
SELECT Medications.name_medication, Medications.description, Medications.price, Medications.inventory,
	Prescriptions.notes, prescriptions_medications.notes
From Medications
INNER JOIN prescriptions_medications on prescriptions_medications.medications_medication_id=Medications.medication_id
INNER JOIN Prescriptions on Prescriptions.prescription_id=prescriptions_medications.prescriptions_prescription_id;

--Asistentele medicale din departmentul de cardiologie si informatii despre departament
SELECT Nurses.last_name, Nurses.first_name,
	Departments.name, Departments.description
From Nurses
INNER JOIN Departments on Departments.department_id=Nurses.department_id
WHERE Departments.department_id=7;

--Pacientii internati in spital in prezent
SELECT Patients.last_name, Patients.first_name, Patients.age,
	Admissions.start_date, Admissions.operations_number, Admissions.status, Admissions.doctor_notes,
	Doctors.last_name, Doctors.first_name
From Patients
INNER JOIN Admissions on Admissions.patient_id=Patients.patient_id
INNER JOIN Doctors on Doctors.doctor_id=Admissions.doctor_id
WHERE Admissions.status='active';

--Examinari efectuate de doctori sefi
SELECT Examinations.examination_date, Examinations.information, Examinations.diagnosis, Examinations.level,
	Doctors.last_name, Doctors.first_name, Doctors.is_Headdoctor
From Examinations
INNER JOIN Appointments on Appointments.appointment_id=Examinations.appointment_id
INNER JOIN Doctors on Doctors.doctor_id=Appointments.doctor_id
WHERE Doctors.is_Headdoctor=TRUE;

--Pacienti si programarile lor
SELECT Patients.first_name, Patients.last_name, Appointments.schedule_date
From Patients
LEFT JOIN Appointments on Patients.patient_id = Appointments.patient_id;

--Doctori si departamentele lor
SELECT Doctors.first_name, Doctors.last_name, Departments.name
From Doctors
LEFT JOIN Departments on Doctors.department_id = Departments.department_id
ORDER By Departments.department_id;

--Internarile si doctorii care le gestioneaza
SELECT Admissions.admission_id, Admissions.status, Doctors.first_name, Doctors.last_name
From Admissions
LEFT JOIN Doctors on Admissions.doctor_id = Doctors.doctor_id;

--Asistentele medicale si departamentele lor
SELECT Nurses.first_name, Nurses.last_name, Departments.name
From Nurses
RIGHT JOIN Departments on Nurses.department_id = Departments.department_id;

--Operatiile si doctorii care le efectueaza
SELECT Doctors.first_name, Doctors.last_name, Operations.information
From Doctors
RIGHT JOIN doctors_operations on doctors_operations.doctors_doctor_id = Doctors.doctor_id
RIGHT JOIN Operations on doctors_operations.operations_operation_id = Operations.operation_id;

--Views

--Operatiile si chirurgii, care le efectueaza
CREATE View ViewOperationsWithDoctors as
SELECT Doctors.first_name, Doctors.last_name, Operations.information
From Doctors
RIGHT JOIN doctors_operations on doctors_operations.doctors_doctor_id = Doctors.doctor_id
RIGHT JOIN Operations on doctors_operations.operations_operation_id = Operations.operation_id;
SELECT * From ViewOperationsWithDoctors;

--Camerele folosite de pacienti internati
CREATE View ViewRoomsUsedByAdmissedPatients as
SELECT Rooms.number_room, Rooms.type_room, Rooms.capacity,
	Admissions.start_date, Admissions.discharge_date, Admissions.status,
	Patients.last_name, Patients.first_name,
	Departments.name
From Rooms
INNER JOIN Admissions on Admissions.room_id=Rooms.room_id
INNER JOIN Patients on Patients.patient_id=Admissions.patient_id
INNER JOIN Departments on Departments.department_id=Rooms.department_id;

--Platile efectuate de pacienti in departamentul de psihiatrie
CREATE View ViewPaymentsInPsychiatry as
SELECT Payment.payment_id, Payment.amount as payment, Payment.date_pay,
	Admissions.operations_number, Admissions.price as admission_price,
	Operations.information, Operations.finished_date, Operations.price as operation_price,
	Patients.last_name, Patients.first_name, Patients.gender
From Payment
LEFT JOIN Operations on Operations.operation_id=Payment.operation_id
LEFT JOIN Admissions on Admissions.admission_id=Operations.admission_id
LEFT JOIN Patients on Patients.patient_id=Admissions.patient_id
LEFT JOIN Doctors on Doctors.doctor_id=Admissions.doctor_id
LEFT JOIN Departments on Departments.department_id=Doctors.department_id
WHERE Departments.name='Psychiatry';
SELECT * From ViewPaymentsInPsychiatry;

--Proceduri

--Programarea pacientilor pentru examinare
CREATE Procedure ProcedurePatientsAppointments(p_patient_id INT, department_name VARCHAR, p_schedule_date TIMESTAMP)
LANGUAGE plpgsql
AS $$
DECLARE
	variable_department_id INT;
	variable_room_id INT;
	variable_doctor_id INT;
BEGIN
	--Verificarea si identificarea numelui de departament
	SELECT department_id
	INTO variable_department_id
	From Departments
	WHERE name=department_name;
	IF NOT FOUND THEN
        RAISE EXCEPTION 'Department "%" does not exist.', department_name;
    END IF;
	
	--Identificarea doctorului pentru programare
	SELECT Doctors.doctor_id
	INTO variable_doctor_id
	From Doctors
	INNER JOIN Departments on Departments.department_id=Doctors.department_id
	WHERE Departments.department_id=variable_department_id and Doctors.is_Headdoctor=TRUE;
	
	--Identificarea camerei pentru examinare
	SELECT Rooms.room_id
	INTO variable_room_id
	From Rooms
	INNER JOIN Departments on Departments.department_id=Rooms.department_id
	WHERE Rooms.type_room='Examination Room' and Departments.department_id=variable_department_id;

	--Inserarea noi instante adaugate de utilizator prin procedura in tabelul de programari
	INSERT Into Appointments (patient_id, doctor_id, room_id, schedule_date, status)
	VALUES (p_patient_id, variable_doctor_id, variable_room_id, p_schedule_date, 'waiting');
END;
$$;
--Verificari
SELECT * From Appointments;
SELECT * From Patients;
SELECT * From Doctors;
--Executia procedurii
CALL ProcedurePatientsAppointments(14, 'Ophthalmology', '2025-05-01 11:35:00');

--Validarea programarii pacientului de catre doctor
CREATE Procedure ProcedureDoctorValidatesAppointment(p_doctor_id INT, p_appointment_id INT, doctor_response VARCHAR)
LANGUAGE plpgsql
AS $$
DECLARE
	variable_doctor_id INT;
BEGIN
	--Verificarea doctorului
	SELECT doctor_id
	INTO variable_doctor_id
	From Appointments
	WHERE appointment_id=p_appointment_id;
	IF NOT FOUND THEN
        RAISE EXCEPTION 'Appointment "%" does not exist.', p_appointment_id;
    END IF;
	IF variable_doctor_id!=p_doctor_id THEN
		RAISE EXCEPTION 'Doctor "%" is not the assigned doctor for the appointment!.', p_doctor_id;
	END IF;
	
	--Actualizarea starii programarii
	IF doctor_response='approved' THEN
		UPDATE Appointments
		SET status = 'scheduled'
		WHERE appointment_id=p_appointment_id;
	ELSIF doctor_response='denied' THEN
		UPDATE Appointments
		SET status='canceled'
		WHERE appointment_id=p_appointment_id;
	ELSE
		RAISE EXCEPTION 'Doctors response "%" is invalid!',doctor_response;
	END IF;
END;
$$;
--Verificari
SELECT * From Appointments;
--Executia procedurii
CALL ProcedureDoctorValidatesAppointment(23, 32, 'approved');

--Examinarea pacientului de catre doctor
CREATE Procedure ProcedureDoctorExaminations(p_doctor_id INT, p_appointment_id INT, p_examination_date TIMESTAMP,
				p_information VARCHAR, p_assistant_nurses TEXT, p_diagnosis VARCHAR, p_result VARCHAR, p_reexamination_date TIMESTAMP,
				p_price DECIMAL, p_level VARCHAR, p_examination_notes TEXT)
LANGUAGE plpgsql
AS $$
DECLARE
	variable_doctor_id INT;
BEGIN
	--Verificarea doctorului
	SELECT doctor_id
	INTO variable_doctor_id
	From Appointments
	WHERE appointment_id=p_appointment_id;
	IF NOT FOUND THEN
        RAISE EXCEPTION 'Appointment "%" does not exist.', p_appointment_id;
    END IF;
	IF variable_doctor_id!=p_doctor_id THEN
		RAISE EXCEPTION 'Doctor "%" is not the assigned doctor for the examination!', p_doctor_id;
	END IF;
	
	--Inserarea instantei in Examinations
	INSERT Into Examinations(appointment_id, examination_date, information, assistant_nurses, diagnosis, result, reexamination_date, price, level, examination_notes)
	VALUES (p_appointment_id, p_examination_date, p_information, p_assistant_nurses, p_diagnosis, p_result, p_reexamination_date, p_price, p_level, p_examination_notes);
END;
$$;
--Verificari
SELECT * From Appointments;
SELECT * From Patients;
SELECT * From Examinations;
--Examinarea
CALL ProcedureDoctorExaminations(23,32,'2025-05-01 12:00:00','Patrick Alexander - eye pain, blurrs','Chloe, Ryan & Lucas','Glaucoma',
	'prescription',null,75.00,'mild','Examination 1: Patient seems to suffer from eye pain and pressure, having developed blind spots.');

--Triggers

--Reducerea automata de medicamente din inventar dupa crearea prescriptiei
--Functia pentru reducerea medicamentelor prescrise pentru pacienti din inventar
CREATE Function ReduceInventoryMedication()
RETURNS Trigger as $$
DECLARE
	current_inventory INT;
BEGIN
	--Inventarul curent prin ID-ul medicamentului din tabela de jonctiune
	SELECT inventory
	INTO current_inventory
	From Medications
	WHERE medication_id=NEW.medications_medication_id;
	
	--Verificarea daca sunt destule medicamente in inventar
	IF current_inventory < 0 THEN
		RAISE EXCEPTION 'Not enough inventory for medication with ID %', NEW.medications_medication_id;
	END IF;
	
	--Reducerea numarului de medicamente din inventar
	UPDATE Medications
	SET inventory = inventory - 1
	WHERE medication_id=NEW.medications_medication_id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
--Triggerul legat de functie care va executa functia dupa inserarea in tabelul de jonctiune cu prescriptie si medicament
CREATE Trigger TriggerReduceInventoryMedication
AFTER INSERT on prescriptions_medications
FOR EACH ROW
EXECUTE FUNCTION ReduceInventoryMedication();
--Verificari
--Prescrierile se fac doar cu un ID din examinari
SELECT * From Examinations;
SELECT * From Prescriptions;
--Verificarea inventarului din tabelul de medicamente
SELECT * From Medications;
--Se adauga inregistrari in tabelul de jonctiune pentru apelul implicit al triggerului, care va reduce inventarul
SELECT * From prescriptions_medications;
--Inserarea
INSERT Into Prescriptions(examination_id, number_medications, notes, prescription_date, total_price)
VALUES (24, 3, 'Prescription: Latanoprost, Timolol and Olopatadine, taken 3 drops in 3 hour difference from each.', '2025-05-01 12:30:00', 199.50);
INSERT Into prescriptions_medications(prescriptions_prescription_id, medications_medication_id, notes)
VALUES (11,7,'Latanoprost prescribed in prescription with ID 6 for patient Patrick Alexander with Glaucoma.'),
(11,8,'Timolol prescribed in prescription with ID 6 for patient Patrick Alexander with Glaucoma.'),
(11,9,'Olopatadine prescribed in prescription with ID 6 for patient Patrick Alexander with Glaucoma.');

--Index

CREATE Index Index_patients_name on Patients(last_name, first_name);
CREATE Index Index_doctors_specialty on Doctors(last_name, first_name, specialty);
CREATE Index Index_nurses_name on Nurses(last_name, first_name);
CREATE Index Index_department_name on Departments(name);
CREATE Index Index_medications_info on Medications(name_medication, price);
CREATE Index Index_rooms_number on Rooms(number_room);
CREATE Index Index_appointments_date on Appointments(patient_id, schedule_date, status);
CREATE Index Index_examinations_result on Examinations(result);
CREATE Index Index_prescriptions_price on Prescriptions(total_price);
CREATE Index Index_admissions_status on Admissions(status);
CREATE Index Index_operations_info on Operations(information,assistant_nurses,assistant_doctors);
CREATE Index Index_payment_amount on Payment(amount);
CREATE Index Index_doctors_operations on doctors_operations(notes);
CREATE Index Index_nurses_doctors on nurses_doctors(notes);
CREATE Index Index_prescriptions_medications on prescriptions_medications(notes);

--Interogari Complexe

--Media numarului de medicamente prescrise in prescriptii
SELECT AVG(number_medications) as media_medicamentelor_din_prescriptii
From Prescriptions;
--Platile totale facute de pacientul cu identificatorul 2
SELECT SUM(amount) as total_payments
From Payment
Where patient_id=2;
--Numarul de programari facute de fiecare pacient
SELECT Patients.first_name, Patients.last_name, COUNT(Appointments.appointment_id) as number_of_appointments
FROM Patients
INNER JOIN Appointments on Patients.patient_id = Appointments.patient_id
GROUP By Patients.patient_id;
--Totalul de plati facute de fiecare pacient mai mari de 1000
SELECT Patients.first_name, Patients.last_name, SUM(Payment.amount) as total_payment
FROM Payment
JOIN Patients on Payment.patient_id = Patients.patient_id
GROUP By Patients.patient_id
HAVING SUM(Payment.amount) > 1000;
