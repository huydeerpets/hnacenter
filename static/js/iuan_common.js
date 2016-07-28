;$(function(){
	//全选的实现
	$(".check-all").click(function(){
		$(".ids").prop("checked", this.checked);
	});
	$(".ids").click(function(){
		var option = $(".ids");
		option.each(function(i){
			if(!this.checked){
				$(".check-all").prop("checked", false);
				return false;
			}else{
				$(".check-all").prop("checked", true);
			}
		});
	});
});

//导航高亮
function highlight_subnav(url){
	alert(2312312)
    //$('.nav-list').find('a[href="'+url+'"]').closest('li').addClass('active');
}