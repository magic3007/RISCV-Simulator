
add.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	1ea50513          	addi	a0,a0,490 # 102a4 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	1a630067          	jr	422(t1) # 10268 <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	cc418193          	addi	gp,gp,-828 # 11d90 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	c4450513          	addi	a0,a0,-956 # 11d18 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	c7460613          	addi	a2,a2,-908 # 11d50 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	25e080e7          	jalr	606(ra) # 10346 <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	1b450513          	addi	a0,a0,436 # 102a4 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	170080e7          	jalr	368(ra) # 10268 <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	1dc080e7          	jalr	476(ra) # 102dc <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	07a080e7          	jalr	122(ra) # 10188 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	16230067          	jr	354(t1) # 10278 <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	bfa7c783          	lbu	a5,-1030(a5) # 11d18 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	43c50513          	addi	a0,a0,1084 # 11570 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	bcf70723          	sb	a5,-1074(a4) # 11d18 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	bcf70023          	sb	a5,-1088(a4) # 11d18 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	bb258593          	addi	a1,a1,-1102 # 11d20 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	3fa50513          	addi	a0,a0,1018 # 11570 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <main>:
int result[10]={1,2,3,4,5,6,7,8,9,10};

//result:11 12 13 14 15 1 2 3 4 5

int main()
{
   10188:	fe010113          	addi	sp,sp,-32
   1018c:	00813c23          	sd	s0,24(sp)
   10190:	02010413          	addi	s0,sp,32
	int i=0;
   10194:	fe042623          	sw	zero,-20(s0)
	for(i=0;i<5;i++)
   10198:	fe042623          	sw	zero,-20(s0)
   1019c:	0480006f          	j	101e4 <main+0x5c>
		result[i]=result[i]+10;
   101a0:	000117b7          	lui	a5,0x11
   101a4:	59078713          	addi	a4,a5,1424 # 11590 <__fini_array_end>
   101a8:	fec42783          	lw	a5,-20(s0)
   101ac:	00279793          	slli	a5,a5,0x2
   101b0:	00f707b3          	add	a5,a4,a5
   101b4:	0007a783          	lw	a5,0(a5)
   101b8:	00a7879b          	addiw	a5,a5,10
   101bc:	0007871b          	sext.w	a4,a5
   101c0:	000117b7          	lui	a5,0x11
   101c4:	59078693          	addi	a3,a5,1424 # 11590 <__fini_array_end>
   101c8:	fec42783          	lw	a5,-20(s0)
   101cc:	00279793          	slli	a5,a5,0x2
   101d0:	00f687b3          	add	a5,a3,a5
   101d4:	00e7a023          	sw	a4,0(a5)
	for(i=0;i<5;i++)
   101d8:	fec42783          	lw	a5,-20(s0)
   101dc:	0017879b          	addiw	a5,a5,1
   101e0:	fef42623          	sw	a5,-20(s0)
   101e4:	fec42783          	lw	a5,-20(s0)
   101e8:	0007871b          	sext.w	a4,a5
   101ec:	00400793          	li	a5,4
   101f0:	fae7d8e3          	bge	a5,a4,101a0 <main+0x18>
	for(i=5;i<10;i++)
   101f4:	00500793          	li	a5,5
   101f8:	fef42623          	sw	a5,-20(s0)
   101fc:	0480006f          	j	10244 <main+0xbc>
		result[i]=result[i]-5;
   10200:	000117b7          	lui	a5,0x11
   10204:	59078713          	addi	a4,a5,1424 # 11590 <__fini_array_end>
   10208:	fec42783          	lw	a5,-20(s0)
   1020c:	00279793          	slli	a5,a5,0x2
   10210:	00f707b3          	add	a5,a4,a5
   10214:	0007a783          	lw	a5,0(a5)
   10218:	ffb7879b          	addiw	a5,a5,-5
   1021c:	0007871b          	sext.w	a4,a5
   10220:	000117b7          	lui	a5,0x11
   10224:	59078693          	addi	a3,a5,1424 # 11590 <__fini_array_end>
   10228:	fec42783          	lw	a5,-20(s0)
   1022c:	00279793          	slli	a5,a5,0x2
   10230:	00f687b3          	add	a5,a3,a5
   10234:	00e7a023          	sw	a4,0(a5)
	for(i=5;i<10;i++)
   10238:	fec42783          	lw	a5,-20(s0)
   1023c:	0017879b          	addiw	a5,a5,1
   10240:	fef42623          	sw	a5,-20(s0)
   10244:	fec42783          	lw	a5,-20(s0)
   10248:	0007871b          	sext.w	a4,a5
   1024c:	00900793          	li	a5,9
   10250:	fae7d8e3          	bge	a5,a4,10200 <main+0x78>
	return 0;
   10254:	00000793          	li	a5,0
   10258:	00078513          	mv	a0,a5
   1025c:	01813403          	ld	s0,24(sp)
   10260:	02010113          	addi	sp,sp,32
   10264:	00008067          	ret

0000000000010268 <atexit>:
   10268:	85aa                	mv	a1,a0
   1026a:	4681                	li	a3,0
   1026c:	4601                	li	a2,0
   1026e:	4501                	li	a0,0
   10270:	00000317          	auipc	t1,0x0
   10274:	18030067          	jr	384(t1) # 103f0 <__register_exitproc>

0000000000010278 <exit>:
   10278:	1141                	addi	sp,sp,-16
   1027a:	4581                	li	a1,0
   1027c:	e022                	sd	s0,0(sp)
   1027e:	e406                	sd	ra,8(sp)
   10280:	842a                	mv	s0,a0
   10282:	00000097          	auipc	ra,0x0
   10286:	1ea080e7          	jalr	490(ra) # 1046c <__call_exitprocs>
   1028a:	00002797          	auipc	a5,0x2
   1028e:	a7678793          	addi	a5,a5,-1418 # 11d00 <_global_impure_ptr>
   10292:	6388                	ld	a0,0(a5)
   10294:	6d3c                	ld	a5,88(a0)
   10296:	c391                	beqz	a5,1029a <exit+0x22>
   10298:	9782                	jalr	a5
   1029a:	8522                	mv	a0,s0
   1029c:	00000097          	auipc	ra,0x0
   102a0:	296080e7          	jalr	662(ra) # 10532 <_exit>

00000000000102a4 <__libc_fini_array>:
   102a4:	1101                	addi	sp,sp,-32
   102a6:	e822                	sd	s0,16(sp)
   102a8:	00001797          	auipc	a5,0x1
   102ac:	2e878793          	addi	a5,a5,744 # 11590 <__fini_array_end>
   102b0:	00001417          	auipc	s0,0x1
   102b4:	2d840413          	addi	s0,s0,728 # 11588 <__init_array_end>
   102b8:	8f81                	sub	a5,a5,s0
   102ba:	e426                	sd	s1,8(sp)
   102bc:	ec06                	sd	ra,24(sp)
   102be:	4037d493          	srai	s1,a5,0x3
   102c2:	c881                	beqz	s1,102d2 <__libc_fini_array+0x2e>
   102c4:	17e1                	addi	a5,a5,-8
   102c6:	943e                	add	s0,s0,a5
   102c8:	601c                	ld	a5,0(s0)
   102ca:	14fd                	addi	s1,s1,-1
   102cc:	1461                	addi	s0,s0,-8
   102ce:	9782                	jalr	a5
   102d0:	fce5                	bnez	s1,102c8 <__libc_fini_array+0x24>
   102d2:	60e2                	ld	ra,24(sp)
   102d4:	6442                	ld	s0,16(sp)
   102d6:	64a2                	ld	s1,8(sp)
   102d8:	6105                	addi	sp,sp,32
   102da:	8082                	ret

00000000000102dc <__libc_init_array>:
   102dc:	1101                	addi	sp,sp,-32
   102de:	e822                	sd	s0,16(sp)
   102e0:	e04a                	sd	s2,0(sp)
   102e2:	00001417          	auipc	s0,0x1
   102e6:	29240413          	addi	s0,s0,658 # 11574 <__preinit_array_end>
   102ea:	00001917          	auipc	s2,0x1
   102ee:	28a90913          	addi	s2,s2,650 # 11574 <__preinit_array_end>
   102f2:	40890933          	sub	s2,s2,s0
   102f6:	ec06                	sd	ra,24(sp)
   102f8:	e426                	sd	s1,8(sp)
   102fa:	40395913          	srai	s2,s2,0x3
   102fe:	00090963          	beqz	s2,10310 <__libc_init_array+0x34>
   10302:	4481                	li	s1,0
   10304:	601c                	ld	a5,0(s0)
   10306:	0485                	addi	s1,s1,1
   10308:	0421                	addi	s0,s0,8
   1030a:	9782                	jalr	a5
   1030c:	fe991ce3          	bne	s2,s1,10304 <__libc_init_array+0x28>
   10310:	00001417          	auipc	s0,0x1
   10314:	26840413          	addi	s0,s0,616 # 11578 <__init_array_start>
   10318:	00001917          	auipc	s2,0x1
   1031c:	27090913          	addi	s2,s2,624 # 11588 <__init_array_end>
   10320:	40890933          	sub	s2,s2,s0
   10324:	40395913          	srai	s2,s2,0x3
   10328:	00090963          	beqz	s2,1033a <__libc_init_array+0x5e>
   1032c:	4481                	li	s1,0
   1032e:	601c                	ld	a5,0(s0)
   10330:	0485                	addi	s1,s1,1
   10332:	0421                	addi	s0,s0,8
   10334:	9782                	jalr	a5
   10336:	fe991ce3          	bne	s2,s1,1032e <__libc_init_array+0x52>
   1033a:	60e2                	ld	ra,24(sp)
   1033c:	6442                	ld	s0,16(sp)
   1033e:	64a2                	ld	s1,8(sp)
   10340:	6902                	ld	s2,0(sp)
   10342:	6105                	addi	sp,sp,32
   10344:	8082                	ret

0000000000010346 <memset>:
   10346:	433d                	li	t1,15
   10348:	872a                	mv	a4,a0
   1034a:	02c37163          	bgeu	t1,a2,1036c <memset+0x26>
   1034e:	00f77793          	andi	a5,a4,15
   10352:	e3c1                	bnez	a5,103d2 <memset+0x8c>
   10354:	e1bd                	bnez	a1,103ba <memset+0x74>
   10356:	ff067693          	andi	a3,a2,-16
   1035a:	8a3d                	andi	a2,a2,15
   1035c:	96ba                	add	a3,a3,a4
   1035e:	e30c                	sd	a1,0(a4)
   10360:	e70c                	sd	a1,8(a4)
   10362:	0741                	addi	a4,a4,16
   10364:	fed76de3          	bltu	a4,a3,1035e <memset+0x18>
   10368:	e211                	bnez	a2,1036c <memset+0x26>
   1036a:	8082                	ret
   1036c:	40c306b3          	sub	a3,t1,a2
   10370:	068a                	slli	a3,a3,0x2
   10372:	00000297          	auipc	t0,0x0
   10376:	9696                	add	a3,a3,t0
   10378:	00a68067          	jr	10(a3)
   1037c:	00b70723          	sb	a1,14(a4)
   10380:	00b706a3          	sb	a1,13(a4)
   10384:	00b70623          	sb	a1,12(a4)
   10388:	00b705a3          	sb	a1,11(a4)
   1038c:	00b70523          	sb	a1,10(a4)
   10390:	00b704a3          	sb	a1,9(a4)
   10394:	00b70423          	sb	a1,8(a4)
   10398:	00b703a3          	sb	a1,7(a4)
   1039c:	00b70323          	sb	a1,6(a4)
   103a0:	00b702a3          	sb	a1,5(a4)
   103a4:	00b70223          	sb	a1,4(a4)
   103a8:	00b701a3          	sb	a1,3(a4)
   103ac:	00b70123          	sb	a1,2(a4)
   103b0:	00b700a3          	sb	a1,1(a4)
   103b4:	00b70023          	sb	a1,0(a4)
   103b8:	8082                	ret
   103ba:	0ff5f593          	andi	a1,a1,255
   103be:	00859693          	slli	a3,a1,0x8
   103c2:	8dd5                	or	a1,a1,a3
   103c4:	01059693          	slli	a3,a1,0x10
   103c8:	8dd5                	or	a1,a1,a3
   103ca:	02059693          	slli	a3,a1,0x20
   103ce:	8dd5                	or	a1,a1,a3
   103d0:	b759                	j	10356 <memset+0x10>
   103d2:	00279693          	slli	a3,a5,0x2
   103d6:	00000297          	auipc	t0,0x0
   103da:	9696                	add	a3,a3,t0
   103dc:	8286                	mv	t0,ra
   103de:	fa2680e7          	jalr	-94(a3)
   103e2:	8096                	mv	ra,t0
   103e4:	17c1                	addi	a5,a5,-16
   103e6:	8f1d                	sub	a4,a4,a5
   103e8:	963e                	add	a2,a2,a5
   103ea:	f8c371e3          	bgeu	t1,a2,1036c <memset+0x26>
   103ee:	b79d                	j	10354 <memset+0xe>

00000000000103f0 <__register_exitproc>:
   103f0:	00002797          	auipc	a5,0x2
   103f4:	91078793          	addi	a5,a5,-1776 # 11d00 <_global_impure_ptr>
   103f8:	6398                	ld	a4,0(a5)
   103fa:	1f873783          	ld	a5,504(a4)
   103fe:	c3b1                	beqz	a5,10442 <__register_exitproc+0x52>
   10400:	4798                	lw	a4,8(a5)
   10402:	487d                	li	a6,31
   10404:	06e84263          	blt	a6,a4,10468 <__register_exitproc+0x78>
   10408:	c505                	beqz	a0,10430 <__register_exitproc+0x40>
   1040a:	00371813          	slli	a6,a4,0x3
   1040e:	983e                	add	a6,a6,a5
   10410:	10c83823          	sd	a2,272(a6)
   10414:	3107a883          	lw	a7,784(a5)
   10418:	4605                	li	a2,1
   1041a:	00e6163b          	sllw	a2,a2,a4
   1041e:	00c8e8b3          	or	a7,a7,a2
   10422:	3117a823          	sw	a7,784(a5)
   10426:	20d83823          	sd	a3,528(a6)
   1042a:	4689                	li	a3,2
   1042c:	02d50063          	beq	a0,a3,1044c <__register_exitproc+0x5c>
   10430:	00270693          	addi	a3,a4,2
   10434:	068e                	slli	a3,a3,0x3
   10436:	2705                	addiw	a4,a4,1
   10438:	c798                	sw	a4,8(a5)
   1043a:	97b6                	add	a5,a5,a3
   1043c:	e38c                	sd	a1,0(a5)
   1043e:	4501                	li	a0,0
   10440:	8082                	ret
   10442:	20070793          	addi	a5,a4,512
   10446:	1ef73c23          	sd	a5,504(a4)
   1044a:	bf5d                	j	10400 <__register_exitproc+0x10>
   1044c:	3147a683          	lw	a3,788(a5)
   10450:	4501                	li	a0,0
   10452:	8e55                	or	a2,a2,a3
   10454:	00270693          	addi	a3,a4,2
   10458:	068e                	slli	a3,a3,0x3
   1045a:	2705                	addiw	a4,a4,1
   1045c:	30c7aa23          	sw	a2,788(a5)
   10460:	c798                	sw	a4,8(a5)
   10462:	97b6                	add	a5,a5,a3
   10464:	e38c                	sd	a1,0(a5)
   10466:	8082                	ret
   10468:	557d                	li	a0,-1
   1046a:	8082                	ret

000000000001046c <__call_exitprocs>:
   1046c:	715d                	addi	sp,sp,-80
   1046e:	00002797          	auipc	a5,0x2
   10472:	89278793          	addi	a5,a5,-1902 # 11d00 <_global_impure_ptr>
   10476:	e062                	sd	s8,0(sp)
   10478:	0007bc03          	ld	s8,0(a5)
   1047c:	f44e                	sd	s3,40(sp)
   1047e:	f052                	sd	s4,32(sp)
   10480:	ec56                	sd	s5,24(sp)
   10482:	e85a                	sd	s6,16(sp)
   10484:	e486                	sd	ra,72(sp)
   10486:	e0a2                	sd	s0,64(sp)
   10488:	fc26                	sd	s1,56(sp)
   1048a:	f84a                	sd	s2,48(sp)
   1048c:	e45e                	sd	s7,8(sp)
   1048e:	8aaa                	mv	s5,a0
   10490:	8b2e                	mv	s6,a1
   10492:	4a05                	li	s4,1
   10494:	59fd                	li	s3,-1
   10496:	1f8c3903          	ld	s2,504(s8)
   1049a:	02090463          	beqz	s2,104c2 <__call_exitprocs+0x56>
   1049e:	00892483          	lw	s1,8(s2)
   104a2:	fff4841b          	addiw	s0,s1,-1
   104a6:	00044e63          	bltz	s0,104c2 <__call_exitprocs+0x56>
   104aa:	048e                	slli	s1,s1,0x3
   104ac:	94ca                	add	s1,s1,s2
   104ae:	020b0663          	beqz	s6,104da <__call_exitprocs+0x6e>
   104b2:	2084b783          	ld	a5,520(s1)
   104b6:	03678263          	beq	a5,s6,104da <__call_exitprocs+0x6e>
   104ba:	347d                	addiw	s0,s0,-1
   104bc:	14e1                	addi	s1,s1,-8
   104be:	ff3418e3          	bne	s0,s3,104ae <__call_exitprocs+0x42>
   104c2:	60a6                	ld	ra,72(sp)
   104c4:	6406                	ld	s0,64(sp)
   104c6:	74e2                	ld	s1,56(sp)
   104c8:	7942                	ld	s2,48(sp)
   104ca:	79a2                	ld	s3,40(sp)
   104cc:	7a02                	ld	s4,32(sp)
   104ce:	6ae2                	ld	s5,24(sp)
   104d0:	6b42                	ld	s6,16(sp)
   104d2:	6ba2                	ld	s7,8(sp)
   104d4:	6c02                	ld	s8,0(sp)
   104d6:	6161                	addi	sp,sp,80
   104d8:	8082                	ret
   104da:	00892783          	lw	a5,8(s2)
   104de:	6498                	ld	a4,8(s1)
   104e0:	37fd                	addiw	a5,a5,-1
   104e2:	04878263          	beq	a5,s0,10526 <__call_exitprocs+0xba>
   104e6:	0004b423          	sd	zero,8(s1)
   104ea:	db61                	beqz	a4,104ba <__call_exitprocs+0x4e>
   104ec:	31092783          	lw	a5,784(s2)
   104f0:	008a16bb          	sllw	a3,s4,s0
   104f4:	00892b83          	lw	s7,8(s2)
   104f8:	8ff5                	and	a5,a5,a3
   104fa:	2781                	sext.w	a5,a5
   104fc:	eb99                	bnez	a5,10512 <__call_exitprocs+0xa6>
   104fe:	9702                	jalr	a4
   10500:	00892783          	lw	a5,8(s2)
   10504:	f97799e3          	bne	a5,s7,10496 <__call_exitprocs+0x2a>
   10508:	1f8c3783          	ld	a5,504(s8)
   1050c:	fb2787e3          	beq	a5,s2,104ba <__call_exitprocs+0x4e>
   10510:	b759                	j	10496 <__call_exitprocs+0x2a>
   10512:	31492783          	lw	a5,788(s2)
   10516:	1084b583          	ld	a1,264(s1)
   1051a:	8ff5                	and	a5,a5,a3
   1051c:	2781                	sext.w	a5,a5
   1051e:	e799                	bnez	a5,1052c <__call_exitprocs+0xc0>
   10520:	8556                	mv	a0,s5
   10522:	9702                	jalr	a4
   10524:	bff1                	j	10500 <__call_exitprocs+0x94>
   10526:	00892423          	sw	s0,8(s2)
   1052a:	b7c1                	j	104ea <__call_exitprocs+0x7e>
   1052c:	852e                	mv	a0,a1
   1052e:	9702                	jalr	a4
   10530:	bfc1                	j	10500 <__call_exitprocs+0x94>

0000000000010532 <_exit>:
   10532:	4581                	li	a1,0
   10534:	4601                	li	a2,0
   10536:	4681                	li	a3,0
   10538:	4701                	li	a4,0
   1053a:	4781                	li	a5,0
   1053c:	05d00893          	li	a7,93
   10540:	00000073          	ecall
   10544:	00054363          	bltz	a0,1054a <_exit+0x18>
   10548:	a001                	j	10548 <_exit+0x16>
   1054a:	1141                	addi	sp,sp,-16
   1054c:	e022                	sd	s0,0(sp)
   1054e:	842a                	mv	s0,a0
   10550:	e406                	sd	ra,8(sp)
   10552:	4080043b          	negw	s0,s0
   10556:	00000097          	auipc	ra,0x0
   1055a:	00c080e7          	jalr	12(ra) # 10562 <__errno>
   1055e:	c100                	sw	s0,0(a0)
   10560:	a001                	j	10560 <_exit+0x2e>

0000000000010562 <__errno>:
   10562:	00001797          	auipc	a5,0x1
   10566:	7ae78793          	addi	a5,a5,1966 # 11d10 <_impure_ptr>
   1056a:	6388                	ld	a0,0(a5)
   1056c:	8082                	ret
