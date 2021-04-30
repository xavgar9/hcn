var myDocuments = [
    [
        {
            GeneralData: {
                ValorationDate: "2020/02/06",
                HCNNumber: "526",
                AdmissionDate: "2020/02/06",
                Room: "26",
            },
            PatientData: {
                FullName: "Benito Antonio Martínez",
                Birthdate: "2020/02/06",
                Gender: "Heterosexual",
                Sex: "M",
                Age: "26", 
                EPS: "PR Salud", 
                Telephone: "31658245", 
                Occupation: "Singer",
                CivilStatus: "Single",                                
            }
        }
    ],

    [
        {
            GeneralData: {
                ValorationDate: "2020/02/26",
                HCNNumber: "44",
                AdmissionDate: "2020/03/26",
                Room: "20",
                Interpretation: "Interpretación datos generales. 10001",
                Feedback: "Feedback datos generales. 10001"
            },
            PatientData: {
                FullName: "José Álvaro Osorio",
                Birthdate: "1985/05/07",
                Gender: "Heterosexual ",
                Sex: "M",
                Age: "26", 
                EPS: "MD Salud", 
                Telephone: "3201452", 
                Occupation: "Singer",
                CivilStatus: "Single",
                Interpretation: "Interpretación de datos del paciente. 10001",
                Feedback: "Feedback PatientData de datos del paciente. 10001"                              
            },
            ConsultationReason: "Razón de consulta dada por el estudiante. 10001",
            Interpretation: "Interpretación general de la HCN. 10001",
            Feedback: "Feedback general de la HCN. 10001"
        },
        {
            GeneralData: {
                ValorationDate: "2020/02/26",
                HCNNumber: "44",
                AdmissionDate: "2020/03/26",
                Room: "20",
                Interpretation: "Interpretación datos generales. 10002",
                Feedback: "Feedback datos generales. 10002"
            },
            PatientData: {
                FullName: "José Álvaro Osorio",
                Birthdate: "1985/05/07",
                Gender: "Heterosexual ",
                Sex: "M",
                Age: "26", 
                EPS: "MD Salud", 
                Telephone: "3201452", 
                Occupation: "Singer",
                CivilStatus: "Single",
                Interpretation: "Interpretación de datos del paciente. 10002",
                Feedback: "Feedback PatientData de datos del paciente. 10002"                              
            },
            ConsultationReason: "Razón de consulta dada por el estudiante. 10002",
            Interpretation: "Interpretación general de la HCN. 10002",
            Feedback: "Feedback general de la HCN. 10002"
        },
        {
            GeneralData: {
                ValorationDate: "2020/02/26",
                HCNNumber: "44",
                AdmissionDate: "2020/03/26",
                Room: "20",
                Feedback: "Feedback datos generales, el estudiante no realizó comentarios. 10003"
            },
            PatientData: {
                FullName: "José Álvaro Osorio",
                Birthdate: "1985/05/07",
                Gender: "Heterosexual ",
                Sex: "M",
                Age: "26", 
                EPS: "MD Salud", 
                Telephone: "3201452", 
                Occupation: "Singer",
                CivilStatus: "Single",
                Interpretation: "Interpretación de datos del paciente. 10003",
                Feedback: "Feedback PatientData de datos del paciente. 10003"                              
            },
            ConsultationReason: "Razón de consulta dada por el estudiante. 10003",
            Interpretation: "Interpretación general de la HCN. 10003",
            Feedback: "Feedback general de la HCN, muy bien, pero recuerde interpretar los datos generales del paciente. 10003"
        },
        {
            GeneralData: {
                ValorationDate: "2020/02/26",
                HCNNumber: "44",
                AdmissionDate: "2020/03/26",
                Room: "20",
                Feedback: "Feedback datos generales. 10004"
            },
            PatientData: {
                FullName: "José Álvaro Osorio",
                Birthdate: "1985/05/07",
                Gender: "Heterosexual ",
                Sex: "M",
                Age: "26", 
                EPS: "MD Salud", 
                Telephone: "3201452", 
                Occupation: "Singer",
                CivilStatus: "Single",
                Feedback: "Feedback PatientData de datos del paciente. 10004"                              
            },
            Feedback: "Feedback general de la HCN. El estudiante no interpretó ningún dato de la HCN. 10004"
        },
        {
            GeneralData: {
                ValorationDate: "2020/02/26",
                HCNNumber: "44",
                AdmissionDate: "2020/03/26",
                Room: "20",
            },
            PatientData: {
                FullName: "José Álvaro Osorio",
                Birthdate: "1985/05/07",
                Gender: "Heterosexual ",
                Sex: "M",
                Age: "26", 
                EPS: "MD Salud", 
                Telephone: "3201452", 
                Occupation: "Singer",
                CivilStatus: "Single",
            }
        }
    ]
]
var mongoIDs = [
    [   // not solved
        ObjectId("607ec7dee81d0518b08d3da0"), // originalHCN
        ObjectId("607ec7dee81d0518b08d3da1"),
        ObjectId("607ec7dee81d0518b08d3da2"),
        ObjectId("607ec7dee81d0518b08d3da3"),
        ObjectId("607ec7dee81d0518b08d3da4"),
        ObjectId("607ec7dee81d0518b08d3da5"),
    ],
    [   // solved
        ObjectId("607ec7dee81d0518b08d3db0"), // originakHCN
        ObjectId("607ec7dee81d0518b08d3db1"),
        ObjectId("607ec7dee81d0518b08d3db2"),
        ObjectId("607ec7dee81d0518b08d3db3"),
        ObjectId("607ec7dee81d0518b08d3db4"),
        ObjectId("607ec7dee81d0518b08d3db5"),
    ]
]
db.HCN.drop()
// documents ready to be used by the students
mongoIDs[0].forEach(id => {
    var document = myDocuments[0][0]
    document["_id"] = id
    db.HCN.insertOne(
        document
    )
})

// documents solved by students
var i
for (i = 0; i < mongoIDs[1].length; i++) {
    var document
    if (i==0){
        document = myDocuments[1][4]    
        document["_id"] = mongoIDs[1][0]    
    } else{
        document = myDocuments[1][i-1]    
        document["_id"] = mongoIDs[1][i]
    }
    db.HCN.insertOne(
        document
    )
}