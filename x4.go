不报错
th:style="${Plan.IsAdoption}==1?'display: none':(${Plan.userid}==${session.user.userID}?'':'display: none' ) "
$lt  10:50:48
报错
th:style="${Plan.userid}!=${session.user.userID}  ?  'display:none':(${Plan.IsAdoption}==1? 'display: none':'')"

th:style="${Plan.userid}!=${session.user.userID}  ?  'display:none':''"