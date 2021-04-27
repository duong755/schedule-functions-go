# Schedule Functions

## API v1

```json
{
    "endPoint": "/api/v1/schedules?studentId={studentId}",
    "method": "*",
    "studentId": {
        "required": true,
        "pattern": "^\\d{8}$"
    },
    "response": [
        {
            "_id": "ObjectId",
            "studentId": "string",
            "studentName": "string",
            "studentBirthday": "string",
            "studentCourse": "string",
            "classId": "string",
            "classNote": "string",
            "studentNote": "string",
            "classes": [
                {
                    "_id": "ObjectId",
                    "subjectId": "string",
                    "subjectName": "string",
                    "credit": "integer",
                    "classId": "string",
                    "teacher": "string",
                    "numberOfStudents": "integer",
                    "session": "string",
                    "weekDay": "integer",
                    "periods": "integer[]",
                    "place": "string",
                    "note": "string"
                }
            ]
        }
    ]
}
```

```json
{
    "endPoint": "/api/v2/classmembers?classId={classId}",
    "method": "*",
    "classId": {
        "required": true
    },
    "response": {
        "_id": "ObjectId",
        "subjectId": "string",
        "subjectName": "string",
        "credit": "integer",
        "classId": "string",
        "teacher": "string",
        "numberOfStudents": "integer",
        "session": "string",
        "weekDay": "integer",
        "periods": "integer[]",
        "place": "string",
        "note": "string",
        "students": [
            {
                "_id": "ObjectId",
                "studentId": "string",
                "studentName": "string",
                "studentNote": "string"
            }
        ]
    }
}
```


## API v2

```json
{
    "endPoint": "/api/v2/schedules?studentId={studentId}",
    "method": "*",
    "studentId": {
        "required": true,
        "pattern": "^\\d{8}$"
    },
    "response": {
        "_id": "ObjectId",
        "studentId": "string",
        "studentName": "string",
        "studentBirthday": "string",
        "studentCourse": "string",
        "classes": [
            {
                "_id": "ObjectId",
                "subjectId": "string",
                "subjectName": "string",
                "credit": "integer",
                "classId": "string",
                "teacher": "string",
                "numberOfStudents": "integer",
                "session": "string",
                "weekDay": "integer",
                "periods": "integer[]",
                "place": "string",
                "note": "string",
                "studentNote": "string"
            }
        ]
    }
}
```

```json
{
    "endPoint": "/api/v2/classmembers?classId={classId}",
    "method": "*",
    "classId": {
        "required": true
    },
    "response": {
        "_id": "ObjectId",
        "subjectId": "string",
        "subjectName": "string",
        "credit": "integer",
        "classId": "string",
        "teacher": "string",
        "numberOfStudents": "integer",
        "session": "string",
        "weekDay": "integer",
        "periods": "integer[]",
        "place": "string",
        "note": "string",
        "students": [
            {
                "_id": "ObjectId",
                "studentId": "string",
                "studentName": "string",
                "studentNote": "string",
                "studentBirthday": "string",
                "studentCourse": "string",
                "classNote": "string"
            }
        ]
    }
}
```
