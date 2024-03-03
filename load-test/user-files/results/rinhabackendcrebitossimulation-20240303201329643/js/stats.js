var stats = {
    type: "GROUP",
name: "All Requests",
path: "",
pathFormatted: "group_missing-name--1146707516",
stats: {
    "name": "All Requests",
    "numberOfRequests": {
        "total": "1963",
        "ok": "1928",
        "ko": "35"
    },
    "minResponseTime": {
        "total": "3",
        "ok": "3",
        "ko": "6"
    },
    "maxResponseTime": {
        "total": "2881",
        "ok": "2881",
        "ko": "105"
    },
    "meanResponseTime": {
        "total": "46",
        "ok": "46",
        "ko": "64"
    },
    "standardDeviation": {
        "total": "287",
        "ok": "289",
        "ko": "29"
    },
    "percentiles1": {
        "total": "6",
        "ok": "6",
        "ko": "72"
    },
    "percentiles2": {
        "total": "8",
        "ok": "8",
        "ko": "86"
    },
    "percentiles3": {
        "total": "27",
        "ok": "12",
        "ko": "100"
    },
    "percentiles4": {
        "total": "2254",
        "ok": "2325",
        "ko": "105"
    },
    "group1": {
    "name": "t < 800 ms",
    "htmlName": "t < 800 ms",
    "count": 1903,
    "percentage": 97
},
    "group2": {
    "name": "800 ms <= t < 1200 ms",
    "htmlName": "t >= 800 ms <br> t < 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group3": {
    "name": "t >= 1200 ms",
    "htmlName": "t >= 1200 ms",
    "count": 25,
    "percentage": 1
},
    "group4": {
    "name": "failed",
    "htmlName": "failed",
    "count": 35,
    "percentage": 2
},
    "meanNumberOfRequestsPerSecond": {
        "total": "7.915",
        "ok": "7.774",
        "ko": "0.141"
    }
},
contents: {
"req_validac-o-es--40002687": {
        type: "REQUEST",
        name: "validações",
path: "validações",
pathFormatted: "req_validac-o-es--40002687",
stats: {
    "name": "validações",
    "numberOfRequests": {
        "total": "103",
        "ok": "68",
        "ko": "35"
    },
    "minResponseTime": {
        "total": "4",
        "ok": "4",
        "ko": "6"
    },
    "maxResponseTime": {
        "total": "2881",
        "ok": "2881",
        "ko": "105"
    },
    "meanResponseTime": {
        "total": "755",
        "ok": "1110",
        "ko": "64"
    },
    "standardDeviation": {
        "total": "1018",
        "ok": "1094",
        "ko": "29"
    },
    "percentiles1": {
        "total": "98",
        "ok": "576",
        "ko": "72"
    },
    "percentiles2": {
        "total": "756",
        "ok": "2525",
        "ko": "86"
    },
    "percentiles3": {
        "total": "2705",
        "ok": "2754",
        "ko": "100"
    },
    "percentiles4": {
        "total": "2779",
        "ok": "2813",
        "ko": "105"
    },
    "group1": {
    "name": "t < 800 ms",
    "htmlName": "t < 800 ms",
    "count": 43,
    "percentage": 42
},
    "group2": {
    "name": "800 ms <= t < 1200 ms",
    "htmlName": "t >= 800 ms <br> t < 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group3": {
    "name": "t >= 1200 ms",
    "htmlName": "t >= 1200 ms",
    "count": 25,
    "percentage": 24
},
    "group4": {
    "name": "failed",
    "htmlName": "failed",
    "count": 35,
    "percentage": 34
},
    "meanNumberOfRequestsPerSecond": {
        "total": "0.415",
        "ok": "0.274",
        "ko": "0.141"
    }
}
    },"req_extratos--1809255608": {
        type: "REQUEST",
        name: "extratos",
path: "extratos",
pathFormatted: "req_extratos--1809255608",
stats: {
    "name": "extratos",
    "numberOfRequests": {
        "total": "1860",
        "ok": "1860",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "3",
        "ok": "3",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "203",
        "ok": "203",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "7",
        "ok": "7",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "9",
        "ok": "9",
        "ko": "-"
    },
    "percentiles1": {
        "total": "6",
        "ok": "6",
        "ko": "-"
    },
    "percentiles2": {
        "total": "8",
        "ok": "8",
        "ko": "-"
    },
    "percentiles3": {
        "total": "10",
        "ok": "10",
        "ko": "-"
    },
    "percentiles4": {
        "total": "21",
        "ok": "21",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "htmlName": "t < 800 ms",
    "count": 1860,
    "percentage": 100
},
    "group2": {
    "name": "800 ms <= t < 1200 ms",
    "htmlName": "t >= 800 ms <br> t < 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group3": {
    "name": "t >= 1200 ms",
    "htmlName": "t >= 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group4": {
    "name": "failed",
    "htmlName": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "7.5",
        "ok": "7.5",
        "ko": "-"
    }
}
    }
}

}

function fillStats(stat){
    $("#numberOfRequests").append(stat.numberOfRequests.total);
    $("#numberOfRequestsOK").append(stat.numberOfRequests.ok);
    $("#numberOfRequestsKO").append(stat.numberOfRequests.ko);

    $("#minResponseTime").append(stat.minResponseTime.total);
    $("#minResponseTimeOK").append(stat.minResponseTime.ok);
    $("#minResponseTimeKO").append(stat.minResponseTime.ko);

    $("#maxResponseTime").append(stat.maxResponseTime.total);
    $("#maxResponseTimeOK").append(stat.maxResponseTime.ok);
    $("#maxResponseTimeKO").append(stat.maxResponseTime.ko);

    $("#meanResponseTime").append(stat.meanResponseTime.total);
    $("#meanResponseTimeOK").append(stat.meanResponseTime.ok);
    $("#meanResponseTimeKO").append(stat.meanResponseTime.ko);

    $("#standardDeviation").append(stat.standardDeviation.total);
    $("#standardDeviationOK").append(stat.standardDeviation.ok);
    $("#standardDeviationKO").append(stat.standardDeviation.ko);

    $("#percentiles1").append(stat.percentiles1.total);
    $("#percentiles1OK").append(stat.percentiles1.ok);
    $("#percentiles1KO").append(stat.percentiles1.ko);

    $("#percentiles2").append(stat.percentiles2.total);
    $("#percentiles2OK").append(stat.percentiles2.ok);
    $("#percentiles2KO").append(stat.percentiles2.ko);

    $("#percentiles3").append(stat.percentiles3.total);
    $("#percentiles3OK").append(stat.percentiles3.ok);
    $("#percentiles3KO").append(stat.percentiles3.ko);

    $("#percentiles4").append(stat.percentiles4.total);
    $("#percentiles4OK").append(stat.percentiles4.ok);
    $("#percentiles4KO").append(stat.percentiles4.ko);

    $("#meanNumberOfRequestsPerSecond").append(stat.meanNumberOfRequestsPerSecond.total);
    $("#meanNumberOfRequestsPerSecondOK").append(stat.meanNumberOfRequestsPerSecond.ok);
    $("#meanNumberOfRequestsPerSecondKO").append(stat.meanNumberOfRequestsPerSecond.ko);
}
