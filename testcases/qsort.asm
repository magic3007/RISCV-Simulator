
qsort.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	36a50513          	addi	a0,a0,874 # 10424 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	32630067          	jr	806(t1) # 103e8 <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	e4418193          	addi	gp,gp,-444 # 11f10 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	d9c50513          	addi	a0,a0,-612 # 11e70 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	e7460613          	addi	a2,a2,-396 # 11f50 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	3de080e7          	jalr	990(ra) # 104c6 <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	33450513          	addi	a0,a0,820 # 10424 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	2f0080e7          	jalr	752(ra) # 103e8 <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	35c080e7          	jalr	860(ra) # 1045c <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	256080e7          	jalr	598(ra) # 10364 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	2e230067          	jr	738(t1) # 103f8 <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	d527c783          	lbu	a5,-686(a5) # 11e70 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	5bc50513          	addi	a0,a0,1468 # 116f0 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	d2f70323          	sb	a5,-730(a4) # 11e70 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	d0f70c23          	sb	a5,-744(a4) # 11e70 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	d0a58593          	addi	a1,a1,-758 # 11e78 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	57a50513          	addi	a0,a0,1402 # 116f0 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <qsort>:
int result[42]={0};

//result: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39

void qsort(int l,int r) 
{  
   10188:	fd010113          	addi	sp,sp,-48
   1018c:	02113423          	sd	ra,40(sp)
   10190:	02813023          	sd	s0,32(sp)
   10194:	03010413          	addi	s0,sp,48
   10198:	00050793          	mv	a5,a0
   1019c:	00058713          	mv	a4,a1
   101a0:	fcf42e23          	sw	a5,-36(s0)
   101a4:	00070793          	mv	a5,a4
   101a8:	fcf42c23          	sw	a5,-40(s0)
	int x = result[l],i = l,j = r;     
   101ac:	000127b7          	lui	a5,0x12
   101b0:	ea878713          	addi	a4,a5,-344 # 11ea8 <result>
   101b4:	fdc42783          	lw	a5,-36(s0)
   101b8:	00279793          	slli	a5,a5,0x2
   101bc:	00f707b3          	add	a5,a4,a5
   101c0:	0007a783          	lw	a5,0(a5)
   101c4:	fef42223          	sw	a5,-28(s0)
   101c8:	fdc42783          	lw	a5,-36(s0)
   101cc:	fef42623          	sw	a5,-20(s0)
   101d0:	fd842783          	lw	a5,-40(s0)
   101d4:	fef42423          	sw	a5,-24(s0)
	if(l >= r) return ;          
   101d8:	fdc42703          	lw	a4,-36(s0)
   101dc:	fd842783          	lw	a5,-40(s0)
   101e0:	0007071b          	sext.w	a4,a4
   101e4:	0007879b          	sext.w	a5,a5
   101e8:	16f75463          	bge	a4,a5,10350 <qsort+0x1c8>
	while(i < j)  
   101ec:	0f00006f          	j	102dc <qsort+0x154>
	{  
		while(i < j && result[j] >= x) j--;    
   101f0:	fe842783          	lw	a5,-24(s0)
   101f4:	fff7879b          	addiw	a5,a5,-1
   101f8:	fef42423          	sw	a5,-24(s0)
   101fc:	fec42703          	lw	a4,-20(s0)
   10200:	fe842783          	lw	a5,-24(s0)
   10204:	0007071b          	sext.w	a4,a4
   10208:	0007879b          	sext.w	a5,a5
   1020c:	02f75463          	bge	a4,a5,10234 <qsort+0xac>
   10210:	000127b7          	lui	a5,0x12
   10214:	ea878713          	addi	a4,a5,-344 # 11ea8 <result>
   10218:	fe842783          	lw	a5,-24(s0)
   1021c:	00279793          	slli	a5,a5,0x2
   10220:	00f707b3          	add	a5,a4,a5
   10224:	0007a703          	lw	a4,0(a5)
   10228:	fe442783          	lw	a5,-28(s0)
   1022c:	0007879b          	sext.w	a5,a5
   10230:	fcf750e3          	bge	a4,a5,101f0 <qsort+0x68>
			result[i] = result[j];                    
   10234:	000127b7          	lui	a5,0x12
   10238:	ea878713          	addi	a4,a5,-344 # 11ea8 <result>
   1023c:	fe842783          	lw	a5,-24(s0)
   10240:	00279793          	slli	a5,a5,0x2
   10244:	00f707b3          	add	a5,a4,a5
   10248:	0007a703          	lw	a4,0(a5)
   1024c:	000127b7          	lui	a5,0x12
   10250:	ea878693          	addi	a3,a5,-344 # 11ea8 <result>
   10254:	fec42783          	lw	a5,-20(s0)
   10258:	00279793          	slli	a5,a5,0x2
   1025c:	00f687b3          	add	a5,a3,a5
   10260:	00e7a023          	sw	a4,0(a5)
		while(i < j && result[i] <= x) i++;       
   10264:	0100006f          	j	10274 <qsort+0xec>
   10268:	fec42783          	lw	a5,-20(s0)
   1026c:	0017879b          	addiw	a5,a5,1
   10270:	fef42623          	sw	a5,-20(s0)
   10274:	fec42703          	lw	a4,-20(s0)
   10278:	fe842783          	lw	a5,-24(s0)
   1027c:	0007071b          	sext.w	a4,a4
   10280:	0007879b          	sext.w	a5,a5
   10284:	02f75463          	bge	a4,a5,102ac <qsort+0x124>
   10288:	000127b7          	lui	a5,0x12
   1028c:	ea878713          	addi	a4,a5,-344 # 11ea8 <result>
   10290:	fec42783          	lw	a5,-20(s0)
   10294:	00279793          	slli	a5,a5,0x2
   10298:	00f707b3          	add	a5,a4,a5
   1029c:	0007a703          	lw	a4,0(a5)
   102a0:	fe442783          	lw	a5,-28(s0)
   102a4:	0007879b          	sext.w	a5,a5
   102a8:	fce7d0e3          	bge	a5,a4,10268 <qsort+0xe0>
			result[j] = result[i];  
   102ac:	000127b7          	lui	a5,0x12
   102b0:	ea878713          	addi	a4,a5,-344 # 11ea8 <result>
   102b4:	fec42783          	lw	a5,-20(s0)
   102b8:	00279793          	slli	a5,a5,0x2
   102bc:	00f707b3          	add	a5,a4,a5
   102c0:	0007a703          	lw	a4,0(a5)
   102c4:	000127b7          	lui	a5,0x12
   102c8:	ea878693          	addi	a3,a5,-344 # 11ea8 <result>
   102cc:	fe842783          	lw	a5,-24(s0)
   102d0:	00279793          	slli	a5,a5,0x2
   102d4:	00f687b3          	add	a5,a3,a5
   102d8:	00e7a023          	sw	a4,0(a5)
	while(i < j)  
   102dc:	fec42703          	lw	a4,-20(s0)
   102e0:	fe842783          	lw	a5,-24(s0)
   102e4:	0007071b          	sext.w	a4,a4
   102e8:	0007879b          	sext.w	a5,a5
   102ec:	f0f748e3          	blt	a4,a5,101fc <qsort+0x74>
	}  
	result[i] = x;                 
   102f0:	000127b7          	lui	a5,0x12
   102f4:	ea878713          	addi	a4,a5,-344 # 11ea8 <result>
   102f8:	fec42783          	lw	a5,-20(s0)
   102fc:	00279793          	slli	a5,a5,0x2
   10300:	00f707b3          	add	a5,a4,a5
   10304:	fe442703          	lw	a4,-28(s0)
   10308:	00e7a023          	sw	a4,0(a5)
	qsort(l,i-1);           
   1030c:	fec42783          	lw	a5,-20(s0)
   10310:	fff7879b          	addiw	a5,a5,-1
   10314:	0007871b          	sext.w	a4,a5
   10318:	fdc42783          	lw	a5,-36(s0)
   1031c:	00070593          	mv	a1,a4
   10320:	00078513          	mv	a0,a5
   10324:	00000097          	auipc	ra,0x0
   10328:	e64080e7          	jalr	-412(ra) # 10188 <qsort>
	qsort(i+1,r);  
   1032c:	fec42783          	lw	a5,-20(s0)
   10330:	0017879b          	addiw	a5,a5,1
   10334:	0007879b          	sext.w	a5,a5
   10338:	fd842703          	lw	a4,-40(s0)
   1033c:	00070593          	mv	a1,a4
   10340:	00078513          	mv	a0,a5
   10344:	00000097          	auipc	ra,0x0
   10348:	e44080e7          	jalr	-444(ra) # 10188 <qsort>
   1034c:	0080006f          	j	10354 <qsort+0x1cc>
	if(l >= r) return ;          
   10350:	00000013          	nop
} 
   10354:	02813083          	ld	ra,40(sp)
   10358:	02013403          	ld	s0,32(sp)
   1035c:	03010113          	addi	sp,sp,48
   10360:	00008067          	ret

0000000000010364 <main>:

int main()  
{
   10364:	fe010113          	addi	sp,sp,-32
   10368:	00113c23          	sd	ra,24(sp)
   1036c:	00813823          	sd	s0,16(sp)
   10370:	02010413          	addi	s0,sp,32
	for(int i = 0; i < 39; i++){
   10374:	fe042623          	sw	zero,-20(s0)
   10378:	0380006f          	j	103b0 <main+0x4c>
		result[i] = 39 - i;
   1037c:	02700713          	li	a4,39
   10380:	fec42783          	lw	a5,-20(s0)
   10384:	40f707bb          	subw	a5,a4,a5
   10388:	0007871b          	sext.w	a4,a5
   1038c:	000127b7          	lui	a5,0x12
   10390:	ea878693          	addi	a3,a5,-344 # 11ea8 <result>
   10394:	fec42783          	lw	a5,-20(s0)
   10398:	00279793          	slli	a5,a5,0x2
   1039c:	00f687b3          	add	a5,a3,a5
   103a0:	00e7a023          	sw	a4,0(a5)
	for(int i = 0; i < 39; i++){
   103a4:	fec42783          	lw	a5,-20(s0)
   103a8:	0017879b          	addiw	a5,a5,1
   103ac:	fef42623          	sw	a5,-20(s0)
   103b0:	fec42783          	lw	a5,-20(s0)
   103b4:	0007871b          	sext.w	a4,a5
   103b8:	02600793          	li	a5,38
   103bc:	fce7d0e3          	bge	a5,a4,1037c <main+0x18>
	}
	qsort(0,39);
   103c0:	02700593          	li	a1,39
   103c4:	00000513          	li	a0,0
   103c8:	00000097          	auipc	ra,0x0
   103cc:	dc0080e7          	jalr	-576(ra) # 10188 <qsort>
	return 0;
   103d0:	00000793          	li	a5,0
   103d4:	00078513          	mv	a0,a5
   103d8:	01813083          	ld	ra,24(sp)
   103dc:	01013403          	ld	s0,16(sp)
   103e0:	02010113          	addi	sp,sp,32
   103e4:	00008067          	ret

00000000000103e8 <atexit>:
   103e8:	85aa                	mv	a1,a0
   103ea:	4681                	li	a3,0
   103ec:	4601                	li	a2,0
   103ee:	4501                	li	a0,0
   103f0:	00000317          	auipc	t1,0x0
   103f4:	18030067          	jr	384(t1) # 10570 <__register_exitproc>

00000000000103f8 <exit>:
   103f8:	1141                	addi	sp,sp,-16
   103fa:	4581                	li	a1,0
   103fc:	e022                	sd	s0,0(sp)
   103fe:	e406                	sd	ra,8(sp)
   10400:	842a                	mv	s0,a0
   10402:	00000097          	auipc	ra,0x0
   10406:	1ea080e7          	jalr	490(ra) # 105ec <__call_exitprocs>
   1040a:	00002797          	auipc	a5,0x2
   1040e:	a4e78793          	addi	a5,a5,-1458 # 11e58 <_global_impure_ptr>
   10412:	6388                	ld	a0,0(a5)
   10414:	6d3c                	ld	a5,88(a0)
   10416:	c391                	beqz	a5,1041a <exit+0x22>
   10418:	9782                	jalr	a5
   1041a:	8522                	mv	a0,s0
   1041c:	00000097          	auipc	ra,0x0
   10420:	296080e7          	jalr	662(ra) # 106b2 <_exit>

0000000000010424 <__libc_fini_array>:
   10424:	1101                	addi	sp,sp,-32
   10426:	e822                	sd	s0,16(sp)
   10428:	00001797          	auipc	a5,0x1
   1042c:	2e878793          	addi	a5,a5,744 # 11710 <__fini_array_end>
   10430:	00001417          	auipc	s0,0x1
   10434:	2d840413          	addi	s0,s0,728 # 11708 <__init_array_end>
   10438:	8f81                	sub	a5,a5,s0
   1043a:	e426                	sd	s1,8(sp)
   1043c:	ec06                	sd	ra,24(sp)
   1043e:	4037d493          	srai	s1,a5,0x3
   10442:	c881                	beqz	s1,10452 <__libc_fini_array+0x2e>
   10444:	17e1                	addi	a5,a5,-8
   10446:	943e                	add	s0,s0,a5
   10448:	601c                	ld	a5,0(s0)
   1044a:	14fd                	addi	s1,s1,-1
   1044c:	1461                	addi	s0,s0,-8
   1044e:	9782                	jalr	a5
   10450:	fce5                	bnez	s1,10448 <__libc_fini_array+0x24>
   10452:	60e2                	ld	ra,24(sp)
   10454:	6442                	ld	s0,16(sp)
   10456:	64a2                	ld	s1,8(sp)
   10458:	6105                	addi	sp,sp,32
   1045a:	8082                	ret

000000000001045c <__libc_init_array>:
   1045c:	1101                	addi	sp,sp,-32
   1045e:	e822                	sd	s0,16(sp)
   10460:	e04a                	sd	s2,0(sp)
   10462:	00001417          	auipc	s0,0x1
   10466:	29240413          	addi	s0,s0,658 # 116f4 <__preinit_array_end>
   1046a:	00001917          	auipc	s2,0x1
   1046e:	28a90913          	addi	s2,s2,650 # 116f4 <__preinit_array_end>
   10472:	40890933          	sub	s2,s2,s0
   10476:	ec06                	sd	ra,24(sp)
   10478:	e426                	sd	s1,8(sp)
   1047a:	40395913          	srai	s2,s2,0x3
   1047e:	00090963          	beqz	s2,10490 <__libc_init_array+0x34>
   10482:	4481                	li	s1,0
   10484:	601c                	ld	a5,0(s0)
   10486:	0485                	addi	s1,s1,1
   10488:	0421                	addi	s0,s0,8
   1048a:	9782                	jalr	a5
   1048c:	fe991ce3          	bne	s2,s1,10484 <__libc_init_array+0x28>
   10490:	00001417          	auipc	s0,0x1
   10494:	26840413          	addi	s0,s0,616 # 116f8 <__init_array_start>
   10498:	00001917          	auipc	s2,0x1
   1049c:	27090913          	addi	s2,s2,624 # 11708 <__init_array_end>
   104a0:	40890933          	sub	s2,s2,s0
   104a4:	40395913          	srai	s2,s2,0x3
   104a8:	00090963          	beqz	s2,104ba <__libc_init_array+0x5e>
   104ac:	4481                	li	s1,0
   104ae:	601c                	ld	a5,0(s0)
   104b0:	0485                	addi	s1,s1,1
   104b2:	0421                	addi	s0,s0,8
   104b4:	9782                	jalr	a5
   104b6:	fe991ce3          	bne	s2,s1,104ae <__libc_init_array+0x52>
   104ba:	60e2                	ld	ra,24(sp)
   104bc:	6442                	ld	s0,16(sp)
   104be:	64a2                	ld	s1,8(sp)
   104c0:	6902                	ld	s2,0(sp)
   104c2:	6105                	addi	sp,sp,32
   104c4:	8082                	ret

00000000000104c6 <memset>:
   104c6:	433d                	li	t1,15
   104c8:	872a                	mv	a4,a0
   104ca:	02c37163          	bgeu	t1,a2,104ec <memset+0x26>
   104ce:	00f77793          	andi	a5,a4,15
   104d2:	e3c1                	bnez	a5,10552 <memset+0x8c>
   104d4:	e1bd                	bnez	a1,1053a <memset+0x74>
   104d6:	ff067693          	andi	a3,a2,-16
   104da:	8a3d                	andi	a2,a2,15
   104dc:	96ba                	add	a3,a3,a4
   104de:	e30c                	sd	a1,0(a4)
   104e0:	e70c                	sd	a1,8(a4)
   104e2:	0741                	addi	a4,a4,16
   104e4:	fed76de3          	bltu	a4,a3,104de <memset+0x18>
   104e8:	e211                	bnez	a2,104ec <memset+0x26>
   104ea:	8082                	ret
   104ec:	40c306b3          	sub	a3,t1,a2
   104f0:	068a                	slli	a3,a3,0x2
   104f2:	00000297          	auipc	t0,0x0
   104f6:	9696                	add	a3,a3,t0
   104f8:	00a68067          	jr	10(a3)
   104fc:	00b70723          	sb	a1,14(a4)
   10500:	00b706a3          	sb	a1,13(a4)
   10504:	00b70623          	sb	a1,12(a4)
   10508:	00b705a3          	sb	a1,11(a4)
   1050c:	00b70523          	sb	a1,10(a4)
   10510:	00b704a3          	sb	a1,9(a4)
   10514:	00b70423          	sb	a1,8(a4)
   10518:	00b703a3          	sb	a1,7(a4)
   1051c:	00b70323          	sb	a1,6(a4)
   10520:	00b702a3          	sb	a1,5(a4)
   10524:	00b70223          	sb	a1,4(a4)
   10528:	00b701a3          	sb	a1,3(a4)
   1052c:	00b70123          	sb	a1,2(a4)
   10530:	00b700a3          	sb	a1,1(a4)
   10534:	00b70023          	sb	a1,0(a4)
   10538:	8082                	ret
   1053a:	0ff5f593          	andi	a1,a1,255
   1053e:	00859693          	slli	a3,a1,0x8
   10542:	8dd5                	or	a1,a1,a3
   10544:	01059693          	slli	a3,a1,0x10
   10548:	8dd5                	or	a1,a1,a3
   1054a:	02059693          	slli	a3,a1,0x20
   1054e:	8dd5                	or	a1,a1,a3
   10550:	b759                	j	104d6 <memset+0x10>
   10552:	00279693          	slli	a3,a5,0x2
   10556:	00000297          	auipc	t0,0x0
   1055a:	9696                	add	a3,a3,t0
   1055c:	8286                	mv	t0,ra
   1055e:	fa2680e7          	jalr	-94(a3)
   10562:	8096                	mv	ra,t0
   10564:	17c1                	addi	a5,a5,-16
   10566:	8f1d                	sub	a4,a4,a5
   10568:	963e                	add	a2,a2,a5
   1056a:	f8c371e3          	bgeu	t1,a2,104ec <memset+0x26>
   1056e:	b79d                	j	104d4 <memset+0xe>

0000000000010570 <__register_exitproc>:
   10570:	00002797          	auipc	a5,0x2
   10574:	8e878793          	addi	a5,a5,-1816 # 11e58 <_global_impure_ptr>
   10578:	6398                	ld	a4,0(a5)
   1057a:	1f873783          	ld	a5,504(a4)
   1057e:	c3b1                	beqz	a5,105c2 <__register_exitproc+0x52>
   10580:	4798                	lw	a4,8(a5)
   10582:	487d                	li	a6,31
   10584:	06e84263          	blt	a6,a4,105e8 <__register_exitproc+0x78>
   10588:	c505                	beqz	a0,105b0 <__register_exitproc+0x40>
   1058a:	00371813          	slli	a6,a4,0x3
   1058e:	983e                	add	a6,a6,a5
   10590:	10c83823          	sd	a2,272(a6)
   10594:	3107a883          	lw	a7,784(a5)
   10598:	4605                	li	a2,1
   1059a:	00e6163b          	sllw	a2,a2,a4
   1059e:	00c8e8b3          	or	a7,a7,a2
   105a2:	3117a823          	sw	a7,784(a5)
   105a6:	20d83823          	sd	a3,528(a6)
   105aa:	4689                	li	a3,2
   105ac:	02d50063          	beq	a0,a3,105cc <__register_exitproc+0x5c>
   105b0:	00270693          	addi	a3,a4,2
   105b4:	068e                	slli	a3,a3,0x3
   105b6:	2705                	addiw	a4,a4,1
   105b8:	c798                	sw	a4,8(a5)
   105ba:	97b6                	add	a5,a5,a3
   105bc:	e38c                	sd	a1,0(a5)
   105be:	4501                	li	a0,0
   105c0:	8082                	ret
   105c2:	20070793          	addi	a5,a4,512
   105c6:	1ef73c23          	sd	a5,504(a4)
   105ca:	bf5d                	j	10580 <__register_exitproc+0x10>
   105cc:	3147a683          	lw	a3,788(a5)
   105d0:	4501                	li	a0,0
   105d2:	8e55                	or	a2,a2,a3
   105d4:	00270693          	addi	a3,a4,2
   105d8:	068e                	slli	a3,a3,0x3
   105da:	2705                	addiw	a4,a4,1
   105dc:	30c7aa23          	sw	a2,788(a5)
   105e0:	c798                	sw	a4,8(a5)
   105e2:	97b6                	add	a5,a5,a3
   105e4:	e38c                	sd	a1,0(a5)
   105e6:	8082                	ret
   105e8:	557d                	li	a0,-1
   105ea:	8082                	ret

00000000000105ec <__call_exitprocs>:
   105ec:	715d                	addi	sp,sp,-80
   105ee:	00002797          	auipc	a5,0x2
   105f2:	86a78793          	addi	a5,a5,-1942 # 11e58 <_global_impure_ptr>
   105f6:	e062                	sd	s8,0(sp)
   105f8:	0007bc03          	ld	s8,0(a5)
   105fc:	f44e                	sd	s3,40(sp)
   105fe:	f052                	sd	s4,32(sp)
   10600:	ec56                	sd	s5,24(sp)
   10602:	e85a                	sd	s6,16(sp)
   10604:	e486                	sd	ra,72(sp)
   10606:	e0a2                	sd	s0,64(sp)
   10608:	fc26                	sd	s1,56(sp)
   1060a:	f84a                	sd	s2,48(sp)
   1060c:	e45e                	sd	s7,8(sp)
   1060e:	8aaa                	mv	s5,a0
   10610:	8b2e                	mv	s6,a1
   10612:	4a05                	li	s4,1
   10614:	59fd                	li	s3,-1
   10616:	1f8c3903          	ld	s2,504(s8)
   1061a:	02090463          	beqz	s2,10642 <__call_exitprocs+0x56>
   1061e:	00892483          	lw	s1,8(s2)
   10622:	fff4841b          	addiw	s0,s1,-1
   10626:	00044e63          	bltz	s0,10642 <__call_exitprocs+0x56>
   1062a:	048e                	slli	s1,s1,0x3
   1062c:	94ca                	add	s1,s1,s2
   1062e:	020b0663          	beqz	s6,1065a <__call_exitprocs+0x6e>
   10632:	2084b783          	ld	a5,520(s1)
   10636:	03678263          	beq	a5,s6,1065a <__call_exitprocs+0x6e>
   1063a:	347d                	addiw	s0,s0,-1
   1063c:	14e1                	addi	s1,s1,-8
   1063e:	ff3418e3          	bne	s0,s3,1062e <__call_exitprocs+0x42>
   10642:	60a6                	ld	ra,72(sp)
   10644:	6406                	ld	s0,64(sp)
   10646:	74e2                	ld	s1,56(sp)
   10648:	7942                	ld	s2,48(sp)
   1064a:	79a2                	ld	s3,40(sp)
   1064c:	7a02                	ld	s4,32(sp)
   1064e:	6ae2                	ld	s5,24(sp)
   10650:	6b42                	ld	s6,16(sp)
   10652:	6ba2                	ld	s7,8(sp)
   10654:	6c02                	ld	s8,0(sp)
   10656:	6161                	addi	sp,sp,80
   10658:	8082                	ret
   1065a:	00892783          	lw	a5,8(s2)
   1065e:	6498                	ld	a4,8(s1)
   10660:	37fd                	addiw	a5,a5,-1
   10662:	04878263          	beq	a5,s0,106a6 <__call_exitprocs+0xba>
   10666:	0004b423          	sd	zero,8(s1)
   1066a:	db61                	beqz	a4,1063a <__call_exitprocs+0x4e>
   1066c:	31092783          	lw	a5,784(s2)
   10670:	008a16bb          	sllw	a3,s4,s0
   10674:	00892b83          	lw	s7,8(s2)
   10678:	8ff5                	and	a5,a5,a3
   1067a:	2781                	sext.w	a5,a5
   1067c:	eb99                	bnez	a5,10692 <__call_exitprocs+0xa6>
   1067e:	9702                	jalr	a4
   10680:	00892783          	lw	a5,8(s2)
   10684:	f97799e3          	bne	a5,s7,10616 <__call_exitprocs+0x2a>
   10688:	1f8c3783          	ld	a5,504(s8)
   1068c:	fb2787e3          	beq	a5,s2,1063a <__call_exitprocs+0x4e>
   10690:	b759                	j	10616 <__call_exitprocs+0x2a>
   10692:	31492783          	lw	a5,788(s2)
   10696:	1084b583          	ld	a1,264(s1)
   1069a:	8ff5                	and	a5,a5,a3
   1069c:	2781                	sext.w	a5,a5
   1069e:	e799                	bnez	a5,106ac <__call_exitprocs+0xc0>
   106a0:	8556                	mv	a0,s5
   106a2:	9702                	jalr	a4
   106a4:	bff1                	j	10680 <__call_exitprocs+0x94>
   106a6:	00892423          	sw	s0,8(s2)
   106aa:	b7c1                	j	1066a <__call_exitprocs+0x7e>
   106ac:	852e                	mv	a0,a1
   106ae:	9702                	jalr	a4
   106b0:	bfc1                	j	10680 <__call_exitprocs+0x94>

00000000000106b2 <_exit>:
   106b2:	4581                	li	a1,0
   106b4:	4601                	li	a2,0
   106b6:	4681                	li	a3,0
   106b8:	4701                	li	a4,0
   106ba:	4781                	li	a5,0
   106bc:	05d00893          	li	a7,93
   106c0:	00000073          	ecall
   106c4:	00054363          	bltz	a0,106ca <_exit+0x18>
   106c8:	a001                	j	106c8 <_exit+0x16>
   106ca:	1141                	addi	sp,sp,-16
   106cc:	e022                	sd	s0,0(sp)
   106ce:	842a                	mv	s0,a0
   106d0:	e406                	sd	ra,8(sp)
   106d2:	4080043b          	negw	s0,s0
   106d6:	00000097          	auipc	ra,0x0
   106da:	00c080e7          	jalr	12(ra) # 106e2 <__errno>
   106de:	c100                	sw	s0,0(a0)
   106e0:	a001                	j	106e0 <_exit+0x2e>

00000000000106e2 <__errno>:
   106e2:	00001797          	auipc	a5,0x1
   106e6:	78678793          	addi	a5,a5,1926 # 11e68 <_impure_ptr>
   106ea:	6388                	ld	a0,0(a5)
   106ec:	8082                	ret
