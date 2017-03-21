</div>
	</body>
<script>
	$('.icon-remove').click(function(){
		var href=$(this).attr('goto');
		var result=window.confirm('确定要这样做吗？');
		if(result){
			window.location.href=href;
		}
	})
</script>
</html>
