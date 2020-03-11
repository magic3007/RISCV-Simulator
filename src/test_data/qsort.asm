
qsort.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	35e50513          	addi	a0,a0,862 # 10418 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	31a30067          	jr	794(t1) # 103dc <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	e3418193          	addi	gp,gp,-460 # 11f00 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	d8c50513          	addi	a0,a0,-628 # 11e60 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	e6460613          	addi	a2,a2,-412 # 11f40 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	3d2080e7          	jalr	978(ra) # 104ba <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	32850513          	addi	a0,a0,808 # 10418 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	2e4080e7          	jalr	740(ra) # 103dc <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	350080e7          	jalr	848(ra) # 10450 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	256080e7          	jalr	598(ra) # 10364 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	2d630067          	jr	726(t1) # 103ec <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	d427c783          	lbu	a5,-702(a5) # 11e60 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	5b050513          	addi	a0,a0,1456 # 116e4 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	d0f70b23          	sb	a5,-746(a4) # 11e60 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	d0f70423          	sb	a5,-760(a4) # 11e60 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	cfa58593          	addi	a1,a1,-774 # 11e68 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	56e50513          	addi	a0,a0,1390 # 116e4 <__FRAME_END__>
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
   101b0:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
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
   10214:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
   10218:	fe842783          	lw	a5,-24(s0)
   1021c:	00279793          	slli	a5,a5,0x2
   10220:	00f707b3          	add	a5,a4,a5
   10224:	0007a703          	lw	a4,0(a5)
   10228:	fe442783          	lw	a5,-28(s0)
   1022c:	0007879b          	sext.w	a5,a5
   10230:	fcf750e3          	bge	a4,a5,101f0 <qsort+0x68>
			result[i] = result[j];                    
   10234:	000127b7          	lui	a5,0x12
   10238:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
   1023c:	fe842783          	lw	a5,-24(s0)
   10240:	00279793          	slli	a5,a5,0x2
   10244:	00f707b3          	add	a5,a4,a5
   10248:	0007a703          	lw	a4,0(a5)
   1024c:	000127b7          	lui	a5,0x12
   10250:	e9878693          	addi	a3,a5,-360 # 11e98 <result>
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
   1028c:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
   10290:	fec42783          	lw	a5,-20(s0)
   10294:	00279793          	slli	a5,a5,0x2
   10298:	00f707b3          	add	a5,a4,a5
   1029c:	0007a703          	lw	a4,0(a5)
   102a0:	fe442783          	lw	a5,-28(s0)
   102a4:	0007879b          	sext.w	a5,a5
   102a8:	fce7d0e3          	bge	a5,a4,10268 <qsort+0xe0>
			result[j] = result[i];  
   102ac:	000127b7          	lui	a5,0x12
   102b0:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
   102b4:	fec42783          	lw	a5,-20(s0)
   102b8:	00279793          	slli	a5,a5,0x2
   102bc:	00f707b3          	add	a5,a4,a5
   102c0:	0007a703          	lw	a4,0(a5)
   102c4:	000127b7          	lui	a5,0x12
   102c8:	e9878693          	addi	a3,a5,-360 # 11e98 <result>
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
   102f4:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
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
	for(int i=40;i>=1;i--)
   10374:	02800793          	li	a5,40
   10378:	fef42623          	sw	a5,-20(s0)
   1037c:	02c0006f          	j	103a8 <main+0x44>
		result[i]=i;
   10380:	000127b7          	lui	a5,0x12
   10384:	e9878713          	addi	a4,a5,-360 # 11e98 <result>
   10388:	fec42783          	lw	a5,-20(s0)
   1038c:	00279793          	slli	a5,a5,0x2
   10390:	00f707b3          	add	a5,a4,a5
   10394:	fec42703          	lw	a4,-20(s0)
   10398:	00e7a023          	sw	a4,0(a5)
	for(int i=40;i>=1;i--)
   1039c:	fec42783          	lw	a5,-20(s0)
   103a0:	fff7879b          	addiw	a5,a5,-1
   103a4:	fef42623          	sw	a5,-20(s0)
   103a8:	fec42783          	lw	a5,-20(s0)
   103ac:	0007879b          	sext.w	a5,a5
   103b0:	fcf048e3          	bgtz	a5,10380 <main+0x1c>
	qsort(0,39);
   103b4:	02700593          	li	a1,39
   103b8:	00000513          	li	a0,0
   103bc:	00000097          	auipc	ra,0x0
   103c0:	dcc080e7          	jalr	-564(ra) # 10188 <qsort>
	return 0;
   103c4:	00000793          	li	a5,0
   103c8:	00078513          	mv	a0,a5
   103cc:	01813083          	ld	ra,24(sp)
   103d0:	01013403          	ld	s0,16(sp)
   103d4:	02010113          	addi	sp,sp,32
   103d8:	00008067          	ret

00000000000103dc <atexit>:
   103dc:	85aa                	mv	a1,a0
   103de:	4681                	li	a3,0
   103e0:	4601                	li	a2,0
   103e2:	4501                	li	a0,0
   103e4:	00000317          	auipc	t1,0x0
   103e8:	18030067          	jr	384(t1) # 10564 <__register_exitproc>

00000000000103ec <exit>:
   103ec:	1141                	addi	sp,sp,-16
   103ee:	4581                	li	a1,0
   103f0:	e022                	sd	s0,0(sp)
   103f2:	e406                	sd	ra,8(sp)
   103f4:	842a                	mv	s0,a0
   103f6:	00000097          	auipc	ra,0x0
   103fa:	1ea080e7          	jalr	490(ra) # 105e0 <__call_exitprocs>
   103fe:	00002797          	auipc	a5,0x2
   10402:	a4a78793          	addi	a5,a5,-1462 # 11e48 <_global_impure_ptr>
   10406:	6388                	ld	a0,0(a5)
   10408:	6d3c                	ld	a5,88(a0)
   1040a:	c391                	beqz	a5,1040e <exit+0x22>
   1040c:	9782                	jalr	a5
   1040e:	8522                	mv	a0,s0
   10410:	00000097          	auipc	ra,0x0
   10414:	296080e7          	jalr	662(ra) # 106a6 <_exit>

0000000000010418 <__libc_fini_array>:
   10418:	1101                	addi	sp,sp,-32
   1041a:	e822                	sd	s0,16(sp)
   1041c:	00001797          	auipc	a5,0x1
   10420:	2e478793          	addi	a5,a5,740 # 11700 <__fini_array_end>
   10424:	00001417          	auipc	s0,0x1
   10428:	2d440413          	addi	s0,s0,724 # 116f8 <__init_array_end>
   1042c:	8f81                	sub	a5,a5,s0
   1042e:	e426                	sd	s1,8(sp)
   10430:	ec06                	sd	ra,24(sp)
   10432:	4037d493          	srai	s1,a5,0x3
   10436:	c881                	beqz	s1,10446 <__libc_fini_array+0x2e>
   10438:	17e1                	addi	a5,a5,-8
   1043a:	943e                	add	s0,s0,a5
   1043c:	601c                	ld	a5,0(s0)
   1043e:	14fd                	addi	s1,s1,-1
   10440:	1461                	addi	s0,s0,-8
   10442:	9782                	jalr	a5
   10444:	fce5                	bnez	s1,1043c <__libc_fini_array+0x24>
   10446:	60e2                	ld	ra,24(sp)
   10448:	6442                	ld	s0,16(sp)
   1044a:	64a2                	ld	s1,8(sp)
   1044c:	6105                	addi	sp,sp,32
   1044e:	8082                	ret

0000000000010450 <__libc_init_array>:
   10450:	1101                	addi	sp,sp,-32
   10452:	e822                	sd	s0,16(sp)
   10454:	e04a                	sd	s2,0(sp)
   10456:	00001417          	auipc	s0,0x1
   1045a:	29240413          	addi	s0,s0,658 # 116e8 <__init_array_start>
   1045e:	00001917          	auipc	s2,0x1
   10462:	28a90913          	addi	s2,s2,650 # 116e8 <__init_array_start>
   10466:	40890933          	sub	s2,s2,s0
   1046a:	ec06                	sd	ra,24(sp)
   1046c:	e426                	sd	s1,8(sp)
   1046e:	40395913          	srai	s2,s2,0x3
   10472:	00090963          	beqz	s2,10484 <__libc_init_array+0x34>
   10476:	4481                	li	s1,0
   10478:	601c                	ld	a5,0(s0)
   1047a:	0485                	addi	s1,s1,1
   1047c:	0421                	addi	s0,s0,8
   1047e:	9782                	jalr	a5
   10480:	fe991ce3          	bne	s2,s1,10478 <__libc_init_array+0x28>
   10484:	00001417          	auipc	s0,0x1
   10488:	26440413          	addi	s0,s0,612 # 116e8 <__init_array_start>
   1048c:	00001917          	auipc	s2,0x1
   10490:	26c90913          	addi	s2,s2,620 # 116f8 <__init_array_end>
   10494:	40890933          	sub	s2,s2,s0
   10498:	40395913          	srai	s2,s2,0x3
   1049c:	00090963          	beqz	s2,104ae <__libc_init_array+0x5e>
   104a0:	4481                	li	s1,0
   104a2:	601c                	ld	a5,0(s0)
   104a4:	0485                	addi	s1,s1,1
   104a6:	0421                	addi	s0,s0,8
   104a8:	9782                	jalr	a5
   104aa:	fe991ce3          	bne	s2,s1,104a2 <__libc_init_array+0x52>
   104ae:	60e2                	ld	ra,24(sp)
   104b0:	6442                	ld	s0,16(sp)
   104b2:	64a2                	ld	s1,8(sp)
   104b4:	6902                	ld	s2,0(sp)
   104b6:	6105                	addi	sp,sp,32
   104b8:	8082                	ret

00000000000104ba <memset>:
   104ba:	433d                	li	t1,15
   104bc:	872a                	mv	a4,a0
   104be:	02c37163          	bgeu	t1,a2,104e0 <memset+0x26>
   104c2:	00f77793          	andi	a5,a4,15
   104c6:	e3c1                	bnez	a5,10546 <memset+0x8c>
   104c8:	e1bd                	bnez	a1,1052e <memset+0x74>
   104ca:	ff067693          	andi	a3,a2,-16
   104ce:	8a3d                	andi	a2,a2,15
   104d0:	96ba                	add	a3,a3,a4
   104d2:	e30c                	sd	a1,0(a4)
   104d4:	e70c                	sd	a1,8(a4)
   104d6:	0741                	addi	a4,a4,16
   104d8:	fed76de3          	bltu	a4,a3,104d2 <memset+0x18>
   104dc:	e211                	bnez	a2,104e0 <memset+0x26>
   104de:	8082                	ret
   104e0:	40c306b3          	sub	a3,t1,a2
   104e4:	068a                	slli	a3,a3,0x2
   104e6:	00000297          	auipc	t0,0x0
   104ea:	9696                	add	a3,a3,t0
   104ec:	00a68067          	jr	10(a3)
   104f0:	00b70723          	sb	a1,14(a4)
   104f4:	00b706a3          	sb	a1,13(a4)
   104f8:	00b70623          	sb	a1,12(a4)
   104fc:	00b705a3          	sb	a1,11(a4)
   10500:	00b70523          	sb	a1,10(a4)
   10504:	00b704a3          	sb	a1,9(a4)
   10508:	00b70423          	sb	a1,8(a4)
   1050c:	00b703a3          	sb	a1,7(a4)
   10510:	00b70323          	sb	a1,6(a4)
   10514:	00b702a3          	sb	a1,5(a4)
   10518:	00b70223          	sb	a1,4(a4)
   1051c:	00b701a3          	sb	a1,3(a4)
   10520:	00b70123          	sb	a1,2(a4)
   10524:	00b700a3          	sb	a1,1(a4)
   10528:	00b70023          	sb	a1,0(a4)
   1052c:	8082                	ret
   1052e:	0ff5f593          	andi	a1,a1,255
   10532:	00859693          	slli	a3,a1,0x8
   10536:	8dd5                	or	a1,a1,a3
   10538:	01059693          	slli	a3,a1,0x10
   1053c:	8dd5                	or	a1,a1,a3
   1053e:	02059693          	slli	a3,a1,0x20
   10542:	8dd5                	or	a1,a1,a3
   10544:	b759                	j	104ca <memset+0x10>
   10546:	00279693          	slli	a3,a5,0x2
   1054a:	00000297          	auipc	t0,0x0
   1054e:	9696                	add	a3,a3,t0
   10550:	8286                	mv	t0,ra
   10552:	fa2680e7          	jalr	-94(a3)
   10556:	8096                	mv	ra,t0
   10558:	17c1                	addi	a5,a5,-16
   1055a:	8f1d                	sub	a4,a4,a5
   1055c:	963e                	add	a2,a2,a5
   1055e:	f8c371e3          	bgeu	t1,a2,104e0 <memset+0x26>
   10562:	b79d                	j	104c8 <memset+0xe>

0000000000010564 <__register_exitproc>:
   10564:	00002797          	auipc	a5,0x2
   10568:	8e478793          	addi	a5,a5,-1820 # 11e48 <_global_impure_ptr>
   1056c:	6398                	ld	a4,0(a5)
   1056e:	1f873783          	ld	a5,504(a4)
   10572:	c3b1                	beqz	a5,105b6 <__register_exitproc+0x52>
   10574:	4798                	lw	a4,8(a5)
   10576:	487d                	li	a6,31
   10578:	06e84263          	blt	a6,a4,105dc <__register_exitproc+0x78>
   1057c:	c505                	beqz	a0,105a4 <__register_exitproc+0x40>
   1057e:	00371813          	slli	a6,a4,0x3
   10582:	983e                	add	a6,a6,a5
   10584:	10c83823          	sd	a2,272(a6)
   10588:	3107a883          	lw	a7,784(a5)
   1058c:	4605                	li	a2,1
   1058e:	00e6163b          	sllw	a2,a2,a4
   10592:	00c8e8b3          	or	a7,a7,a2
   10596:	3117a823          	sw	a7,784(a5)
   1059a:	20d83823          	sd	a3,528(a6)
   1059e:	4689                	li	a3,2
   105a0:	02d50063          	beq	a0,a3,105c0 <__register_exitproc+0x5c>
   105a4:	00270693          	addi	a3,a4,2
   105a8:	068e                	slli	a3,a3,0x3
   105aa:	2705                	addiw	a4,a4,1
   105ac:	c798                	sw	a4,8(a5)
   105ae:	97b6                	add	a5,a5,a3
   105b0:	e38c                	sd	a1,0(a5)
   105b2:	4501                	li	a0,0
   105b4:	8082                	ret
   105b6:	20070793          	addi	a5,a4,512
   105ba:	1ef73c23          	sd	a5,504(a4)
   105be:	bf5d                	j	10574 <__register_exitproc+0x10>
   105c0:	3147a683          	lw	a3,788(a5)
   105c4:	4501                	li	a0,0
   105c6:	8e55                	or	a2,a2,a3
   105c8:	00270693          	addi	a3,a4,2
   105cc:	068e                	slli	a3,a3,0x3
   105ce:	2705                	addiw	a4,a4,1
   105d0:	30c7aa23          	sw	a2,788(a5)
   105d4:	c798                	sw	a4,8(a5)
   105d6:	97b6                	add	a5,a5,a3
   105d8:	e38c                	sd	a1,0(a5)
   105da:	8082                	ret
   105dc:	557d                	li	a0,-1
   105de:	8082                	ret

00000000000105e0 <__call_exitprocs>:
   105e0:	715d                	addi	sp,sp,-80
   105e2:	00002797          	auipc	a5,0x2
   105e6:	86678793          	addi	a5,a5,-1946 # 11e48 <_global_impure_ptr>
   105ea:	e062                	sd	s8,0(sp)
   105ec:	0007bc03          	ld	s8,0(a5)
   105f0:	f44e                	sd	s3,40(sp)
   105f2:	f052                	sd	s4,32(sp)
   105f4:	ec56                	sd	s5,24(sp)
   105f6:	e85a                	sd	s6,16(sp)
   105f8:	e486                	sd	ra,72(sp)
   105fa:	e0a2                	sd	s0,64(sp)
   105fc:	fc26                	sd	s1,56(sp)
   105fe:	f84a                	sd	s2,48(sp)
   10600:	e45e                	sd	s7,8(sp)
   10602:	8aaa                	mv	s5,a0
   10604:	8b2e                	mv	s6,a1
   10606:	4a05                	li	s4,1
   10608:	59fd                	li	s3,-1
   1060a:	1f8c3903          	ld	s2,504(s8)
   1060e:	02090463          	beqz	s2,10636 <__call_exitprocs+0x56>
   10612:	00892483          	lw	s1,8(s2)
   10616:	fff4841b          	addiw	s0,s1,-1
   1061a:	00044e63          	bltz	s0,10636 <__call_exitprocs+0x56>
   1061e:	048e                	slli	s1,s1,0x3
   10620:	94ca                	add	s1,s1,s2
   10622:	020b0663          	beqz	s6,1064e <__call_exitprocs+0x6e>
   10626:	2084b783          	ld	a5,520(s1)
   1062a:	03678263          	beq	a5,s6,1064e <__call_exitprocs+0x6e>
   1062e:	347d                	addiw	s0,s0,-1
   10630:	14e1                	addi	s1,s1,-8
   10632:	ff3418e3          	bne	s0,s3,10622 <__call_exitprocs+0x42>
   10636:	60a6                	ld	ra,72(sp)
   10638:	6406                	ld	s0,64(sp)
   1063a:	74e2                	ld	s1,56(sp)
   1063c:	7942                	ld	s2,48(sp)
   1063e:	79a2                	ld	s3,40(sp)
   10640:	7a02                	ld	s4,32(sp)
   10642:	6ae2                	ld	s5,24(sp)
   10644:	6b42                	ld	s6,16(sp)
   10646:	6ba2                	ld	s7,8(sp)
   10648:	6c02                	ld	s8,0(sp)
   1064a:	6161                	addi	sp,sp,80
   1064c:	8082                	ret
   1064e:	00892783          	lw	a5,8(s2)
   10652:	6498                	ld	a4,8(s1)
   10654:	37fd                	addiw	a5,a5,-1
   10656:	04878263          	beq	a5,s0,1069a <__call_exitprocs+0xba>
   1065a:	0004b423          	sd	zero,8(s1)
   1065e:	db61                	beqz	a4,1062e <__call_exitprocs+0x4e>
   10660:	31092783          	lw	a5,784(s2)
   10664:	008a16bb          	sllw	a3,s4,s0
   10668:	00892b83          	lw	s7,8(s2)
   1066c:	8ff5                	and	a5,a5,a3
   1066e:	2781                	sext.w	a5,a5
   10670:	eb99                	bnez	a5,10686 <__call_exitprocs+0xa6>
   10672:	9702                	jalr	a4
   10674:	00892783          	lw	a5,8(s2)
   10678:	f97799e3          	bne	a5,s7,1060a <__call_exitprocs+0x2a>
   1067c:	1f8c3783          	ld	a5,504(s8)
   10680:	fb2787e3          	beq	a5,s2,1062e <__call_exitprocs+0x4e>
   10684:	b759                	j	1060a <__call_exitprocs+0x2a>
   10686:	31492783          	lw	a5,788(s2)
   1068a:	1084b583          	ld	a1,264(s1)
   1068e:	8ff5                	and	a5,a5,a3
   10690:	2781                	sext.w	a5,a5
   10692:	e799                	bnez	a5,106a0 <__call_exitprocs+0xc0>
   10694:	8556                	mv	a0,s5
   10696:	9702                	jalr	a4
   10698:	bff1                	j	10674 <__call_exitprocs+0x94>
   1069a:	00892423          	sw	s0,8(s2)
   1069e:	b7c1                	j	1065e <__call_exitprocs+0x7e>
   106a0:	852e                	mv	a0,a1
   106a2:	9702                	jalr	a4
   106a4:	bfc1                	j	10674 <__call_exitprocs+0x94>

00000000000106a6 <_exit>:
   106a6:	4581                	li	a1,0
   106a8:	4601                	li	a2,0
   106aa:	4681                	li	a3,0
   106ac:	4701                	li	a4,0
   106ae:	4781                	li	a5,0
   106b0:	05d00893          	li	a7,93
   106b4:	00000073          	ecall
   106b8:	00054363          	bltz	a0,106be <_exit+0x18>
   106bc:	a001                	j	106bc <_exit+0x16>
   106be:	1141                	addi	sp,sp,-16
   106c0:	e022                	sd	s0,0(sp)
   106c2:	842a                	mv	s0,a0
   106c4:	e406                	sd	ra,8(sp)
   106c6:	4080043b          	negw	s0,s0
   106ca:	00000097          	auipc	ra,0x0
   106ce:	00c080e7          	jalr	12(ra) # 106d6 <__errno>
   106d2:	c100                	sw	s0,0(a0)
   106d4:	a001                	j	106d4 <_exit+0x2e>

00000000000106d6 <__errno>:
   106d6:	00001797          	auipc	a5,0x1
   106da:	78278793          	addi	a5,a5,1922 # 11e58 <_impure_ptr>
   106de:	6388                	ld	a0,0(a5)
   106e0:	8082                	ret
