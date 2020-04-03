
mul-div.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	1fe50513          	addi	a0,a0,510 # 102b8 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	1ba30067          	jr	442(t1) # 1027c <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	cd418193          	addi	gp,gp,-812 # 11da0 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	c5450513          	addi	a0,a0,-940 # 11d28 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	c8460613          	addi	a2,a2,-892 # 11d60 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	272080e7          	jalr	626(ra) # 1035a <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	1c850513          	addi	a0,a0,456 # 102b8 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	184080e7          	jalr	388(ra) # 1027c <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	1f0080e7          	jalr	496(ra) # 102f0 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	07a080e7          	jalr	122(ra) # 10188 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	17630067          	jr	374(t1) # 1028c <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	c0a7c783          	lbu	a5,-1014(a5) # 11d28 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	45050513          	addi	a0,a0,1104 # 11584 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	bcf70f23          	sb	a5,-1058(a4) # 11d28 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	bcf70823          	sb	a5,-1072(a4) # 11d28 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	bc258593          	addi	a1,a1,-1086 # 11d30 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	40e50513          	addi	a0,a0,1038 # 11584 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <main>:
int result[10]={1,2,3,4,5,2,4,6,8,10};

//result:5 10 15 20 25 1 2 3 4 5

int main()
{
   10188:	fe010113          	addi	sp,sp,-32
   1018c:	00813c23          	sd	s0,24(sp)
   10190:	02010413          	addi	s0,sp,32
	int i=0;
   10194:	fe042623          	sw	zero,-20(s0)
	for(i=0;i<5;i++)
   10198:	fe042623          	sw	zero,-20(s0)
   1019c:	0540006f          	j	101f0 <main+0x68>
		result[i]=result[i]*5;
   101a0:	000117b7          	lui	a5,0x11
   101a4:	5a078713          	addi	a4,a5,1440 # 115a0 <__fini_array_end>
   101a8:	fec42783          	lw	a5,-20(s0)
   101ac:	00279793          	slli	a5,a5,0x2
   101b0:	00f707b3          	add	a5,a4,a5
   101b4:	0007a783          	lw	a5,0(a5)
   101b8:	00078713          	mv	a4,a5
   101bc:	00070793          	mv	a5,a4
   101c0:	0027979b          	slliw	a5,a5,0x2
   101c4:	00e787bb          	addw	a5,a5,a4
   101c8:	0007871b          	sext.w	a4,a5
   101cc:	000117b7          	lui	a5,0x11
   101d0:	5a078693          	addi	a3,a5,1440 # 115a0 <__fini_array_end>
   101d4:	fec42783          	lw	a5,-20(s0)
   101d8:	00279793          	slli	a5,a5,0x2
   101dc:	00f687b3          	add	a5,a3,a5
   101e0:	00e7a023          	sw	a4,0(a5)
	for(i=0;i<5;i++)
   101e4:	fec42783          	lw	a5,-20(s0)
   101e8:	0017879b          	addiw	a5,a5,1
   101ec:	fef42623          	sw	a5,-20(s0)
   101f0:	fec42783          	lw	a5,-20(s0)
   101f4:	0007871b          	sext.w	a4,a5
   101f8:	00400793          	li	a5,4
   101fc:	fae7d2e3          	bge	a5,a4,101a0 <main+0x18>
	for(i=5;i<10;i++)
   10200:	00500793          	li	a5,5
   10204:	fef42623          	sw	a5,-20(s0)
   10208:	0500006f          	j	10258 <main+0xd0>
		result[i]=result[i]/2;
   1020c:	000117b7          	lui	a5,0x11
   10210:	5a078713          	addi	a4,a5,1440 # 115a0 <__fini_array_end>
   10214:	fec42783          	lw	a5,-20(s0)
   10218:	00279793          	slli	a5,a5,0x2
   1021c:	00f707b3          	add	a5,a4,a5
   10220:	0007a783          	lw	a5,0(a5)
   10224:	01f7d71b          	srliw	a4,a5,0x1f
   10228:	00f707bb          	addw	a5,a4,a5
   1022c:	4017d79b          	sraiw	a5,a5,0x1
   10230:	0007871b          	sext.w	a4,a5
   10234:	000117b7          	lui	a5,0x11
   10238:	5a078693          	addi	a3,a5,1440 # 115a0 <__fini_array_end>
   1023c:	fec42783          	lw	a5,-20(s0)
   10240:	00279793          	slli	a5,a5,0x2
   10244:	00f687b3          	add	a5,a3,a5
   10248:	00e7a023          	sw	a4,0(a5)
	for(i=5;i<10;i++)
   1024c:	fec42783          	lw	a5,-20(s0)
   10250:	0017879b          	addiw	a5,a5,1
   10254:	fef42623          	sw	a5,-20(s0)
   10258:	fec42783          	lw	a5,-20(s0)
   1025c:	0007871b          	sext.w	a4,a5
   10260:	00900793          	li	a5,9
   10264:	fae7d4e3          	bge	a5,a4,1020c <main+0x84>
	return 0;
   10268:	00000793          	li	a5,0
   1026c:	00078513          	mv	a0,a5
   10270:	01813403          	ld	s0,24(sp)
   10274:	02010113          	addi	sp,sp,32
   10278:	00008067          	ret

000000000001027c <atexit>:
   1027c:	85aa                	mv	a1,a0
   1027e:	4681                	li	a3,0
   10280:	4601                	li	a2,0
   10282:	4501                	li	a0,0
   10284:	00000317          	auipc	t1,0x0
   10288:	18030067          	jr	384(t1) # 10404 <__register_exitproc>

000000000001028c <exit>:
   1028c:	1141                	addi	sp,sp,-16
   1028e:	4581                	li	a1,0
   10290:	e022                	sd	s0,0(sp)
   10292:	e406                	sd	ra,8(sp)
   10294:	842a                	mv	s0,a0
   10296:	00000097          	auipc	ra,0x0
   1029a:	1ea080e7          	jalr	490(ra) # 10480 <__call_exitprocs>
   1029e:	00002797          	auipc	a5,0x2
   102a2:	a7278793          	addi	a5,a5,-1422 # 11d10 <_global_impure_ptr>
   102a6:	6388                	ld	a0,0(a5)
   102a8:	6d3c                	ld	a5,88(a0)
   102aa:	c391                	beqz	a5,102ae <exit+0x22>
   102ac:	9782                	jalr	a5
   102ae:	8522                	mv	a0,s0
   102b0:	00000097          	auipc	ra,0x0
   102b4:	296080e7          	jalr	662(ra) # 10546 <_exit>

00000000000102b8 <__libc_fini_array>:
   102b8:	1101                	addi	sp,sp,-32
   102ba:	e822                	sd	s0,16(sp)
   102bc:	00001797          	auipc	a5,0x1
   102c0:	2e478793          	addi	a5,a5,740 # 115a0 <__fini_array_end>
   102c4:	00001417          	auipc	s0,0x1
   102c8:	2d440413          	addi	s0,s0,724 # 11598 <__init_array_end>
   102cc:	8f81                	sub	a5,a5,s0
   102ce:	e426                	sd	s1,8(sp)
   102d0:	ec06                	sd	ra,24(sp)
   102d2:	4037d493          	srai	s1,a5,0x3
   102d6:	c881                	beqz	s1,102e6 <__libc_fini_array+0x2e>
   102d8:	17e1                	addi	a5,a5,-8
   102da:	943e                	add	s0,s0,a5
   102dc:	601c                	ld	a5,0(s0)
   102de:	14fd                	addi	s1,s1,-1
   102e0:	1461                	addi	s0,s0,-8
   102e2:	9782                	jalr	a5
   102e4:	fce5                	bnez	s1,102dc <__libc_fini_array+0x24>
   102e6:	60e2                	ld	ra,24(sp)
   102e8:	6442                	ld	s0,16(sp)
   102ea:	64a2                	ld	s1,8(sp)
   102ec:	6105                	addi	sp,sp,32
   102ee:	8082                	ret

00000000000102f0 <__libc_init_array>:
   102f0:	1101                	addi	sp,sp,-32
   102f2:	e822                	sd	s0,16(sp)
   102f4:	e04a                	sd	s2,0(sp)
   102f6:	00001417          	auipc	s0,0x1
   102fa:	29240413          	addi	s0,s0,658 # 11588 <__init_array_start>
   102fe:	00001917          	auipc	s2,0x1
   10302:	28a90913          	addi	s2,s2,650 # 11588 <__init_array_start>
   10306:	40890933          	sub	s2,s2,s0
   1030a:	ec06                	sd	ra,24(sp)
   1030c:	e426                	sd	s1,8(sp)
   1030e:	40395913          	srai	s2,s2,0x3
   10312:	00090963          	beqz	s2,10324 <__libc_init_array+0x34>
   10316:	4481                	li	s1,0
   10318:	601c                	ld	a5,0(s0)
   1031a:	0485                	addi	s1,s1,1
   1031c:	0421                	addi	s0,s0,8
   1031e:	9782                	jalr	a5
   10320:	fe991ce3          	bne	s2,s1,10318 <__libc_init_array+0x28>
   10324:	00001417          	auipc	s0,0x1
   10328:	26440413          	addi	s0,s0,612 # 11588 <__init_array_start>
   1032c:	00001917          	auipc	s2,0x1
   10330:	26c90913          	addi	s2,s2,620 # 11598 <__init_array_end>
   10334:	40890933          	sub	s2,s2,s0
   10338:	40395913          	srai	s2,s2,0x3
   1033c:	00090963          	beqz	s2,1034e <__libc_init_array+0x5e>
   10340:	4481                	li	s1,0
   10342:	601c                	ld	a5,0(s0)
   10344:	0485                	addi	s1,s1,1
   10346:	0421                	addi	s0,s0,8
   10348:	9782                	jalr	a5
   1034a:	fe991ce3          	bne	s2,s1,10342 <__libc_init_array+0x52>
   1034e:	60e2                	ld	ra,24(sp)
   10350:	6442                	ld	s0,16(sp)
   10352:	64a2                	ld	s1,8(sp)
   10354:	6902                	ld	s2,0(sp)
   10356:	6105                	addi	sp,sp,32
   10358:	8082                	ret

000000000001035a <memset>:
   1035a:	433d                	li	t1,15
   1035c:	872a                	mv	a4,a0
   1035e:	02c37163          	bgeu	t1,a2,10380 <memset+0x26>
   10362:	00f77793          	andi	a5,a4,15
   10366:	e3c1                	bnez	a5,103e6 <memset+0x8c>
   10368:	e1bd                	bnez	a1,103ce <memset+0x74>
   1036a:	ff067693          	andi	a3,a2,-16
   1036e:	8a3d                	andi	a2,a2,15
   10370:	96ba                	add	a3,a3,a4
   10372:	e30c                	sd	a1,0(a4)
   10374:	e70c                	sd	a1,8(a4)
   10376:	0741                	addi	a4,a4,16
   10378:	fed76de3          	bltu	a4,a3,10372 <memset+0x18>
   1037c:	e211                	bnez	a2,10380 <memset+0x26>
   1037e:	8082                	ret
   10380:	40c306b3          	sub	a3,t1,a2
   10384:	068a                	slli	a3,a3,0x2
   10386:	00000297          	auipc	t0,0x0
   1038a:	9696                	add	a3,a3,t0
   1038c:	00a68067          	jr	10(a3)
   10390:	00b70723          	sb	a1,14(a4)
   10394:	00b706a3          	sb	a1,13(a4)
   10398:	00b70623          	sb	a1,12(a4)
   1039c:	00b705a3          	sb	a1,11(a4)
   103a0:	00b70523          	sb	a1,10(a4)
   103a4:	00b704a3          	sb	a1,9(a4)
   103a8:	00b70423          	sb	a1,8(a4)
   103ac:	00b703a3          	sb	a1,7(a4)
   103b0:	00b70323          	sb	a1,6(a4)
   103b4:	00b702a3          	sb	a1,5(a4)
   103b8:	00b70223          	sb	a1,4(a4)
   103bc:	00b701a3          	sb	a1,3(a4)
   103c0:	00b70123          	sb	a1,2(a4)
   103c4:	00b700a3          	sb	a1,1(a4)
   103c8:	00b70023          	sb	a1,0(a4)
   103cc:	8082                	ret
   103ce:	0ff5f593          	andi	a1,a1,255
   103d2:	00859693          	slli	a3,a1,0x8
   103d6:	8dd5                	or	a1,a1,a3
   103d8:	01059693          	slli	a3,a1,0x10
   103dc:	8dd5                	or	a1,a1,a3
   103de:	02059693          	slli	a3,a1,0x20
   103e2:	8dd5                	or	a1,a1,a3
   103e4:	b759                	j	1036a <memset+0x10>
   103e6:	00279693          	slli	a3,a5,0x2
   103ea:	00000297          	auipc	t0,0x0
   103ee:	9696                	add	a3,a3,t0
   103f0:	8286                	mv	t0,ra
   103f2:	fa2680e7          	jalr	-94(a3)
   103f6:	8096                	mv	ra,t0
   103f8:	17c1                	addi	a5,a5,-16
   103fa:	8f1d                	sub	a4,a4,a5
   103fc:	963e                	add	a2,a2,a5
   103fe:	f8c371e3          	bgeu	t1,a2,10380 <memset+0x26>
   10402:	b79d                	j	10368 <memset+0xe>

0000000000010404 <__register_exitproc>:
   10404:	00002797          	auipc	a5,0x2
   10408:	90c78793          	addi	a5,a5,-1780 # 11d10 <_global_impure_ptr>
   1040c:	6398                	ld	a4,0(a5)
   1040e:	1f873783          	ld	a5,504(a4)
   10412:	c3b1                	beqz	a5,10456 <__register_exitproc+0x52>
   10414:	4798                	lw	a4,8(a5)
   10416:	487d                	li	a6,31
   10418:	06e84263          	blt	a6,a4,1047c <__register_exitproc+0x78>
   1041c:	c505                	beqz	a0,10444 <__register_exitproc+0x40>
   1041e:	00371813          	slli	a6,a4,0x3
   10422:	983e                	add	a6,a6,a5
   10424:	10c83823          	sd	a2,272(a6)
   10428:	3107a883          	lw	a7,784(a5)
   1042c:	4605                	li	a2,1
   1042e:	00e6163b          	sllw	a2,a2,a4
   10432:	00c8e8b3          	or	a7,a7,a2
   10436:	3117a823          	sw	a7,784(a5)
   1043a:	20d83823          	sd	a3,528(a6)
   1043e:	4689                	li	a3,2
   10440:	02d50063          	beq	a0,a3,10460 <__register_exitproc+0x5c>
   10444:	00270693          	addi	a3,a4,2
   10448:	068e                	slli	a3,a3,0x3
   1044a:	2705                	addiw	a4,a4,1
   1044c:	c798                	sw	a4,8(a5)
   1044e:	97b6                	add	a5,a5,a3
   10450:	e38c                	sd	a1,0(a5)
   10452:	4501                	li	a0,0
   10454:	8082                	ret
   10456:	20070793          	addi	a5,a4,512
   1045a:	1ef73c23          	sd	a5,504(a4)
   1045e:	bf5d                	j	10414 <__register_exitproc+0x10>
   10460:	3147a683          	lw	a3,788(a5)
   10464:	4501                	li	a0,0
   10466:	8e55                	or	a2,a2,a3
   10468:	00270693          	addi	a3,a4,2
   1046c:	068e                	slli	a3,a3,0x3
   1046e:	2705                	addiw	a4,a4,1
   10470:	30c7aa23          	sw	a2,788(a5)
   10474:	c798                	sw	a4,8(a5)
   10476:	97b6                	add	a5,a5,a3
   10478:	e38c                	sd	a1,0(a5)
   1047a:	8082                	ret
   1047c:	557d                	li	a0,-1
   1047e:	8082                	ret

0000000000010480 <__call_exitprocs>:
   10480:	715d                	addi	sp,sp,-80
   10482:	00002797          	auipc	a5,0x2
   10486:	88e78793          	addi	a5,a5,-1906 # 11d10 <_global_impure_ptr>
   1048a:	e062                	sd	s8,0(sp)
   1048c:	0007bc03          	ld	s8,0(a5)
   10490:	f44e                	sd	s3,40(sp)
   10492:	f052                	sd	s4,32(sp)
   10494:	ec56                	sd	s5,24(sp)
   10496:	e85a                	sd	s6,16(sp)
   10498:	e486                	sd	ra,72(sp)
   1049a:	e0a2                	sd	s0,64(sp)
   1049c:	fc26                	sd	s1,56(sp)
   1049e:	f84a                	sd	s2,48(sp)
   104a0:	e45e                	sd	s7,8(sp)
   104a2:	8aaa                	mv	s5,a0
   104a4:	8b2e                	mv	s6,a1
   104a6:	4a05                	li	s4,1
   104a8:	59fd                	li	s3,-1
   104aa:	1f8c3903          	ld	s2,504(s8)
   104ae:	02090463          	beqz	s2,104d6 <__call_exitprocs+0x56>
   104b2:	00892483          	lw	s1,8(s2)
   104b6:	fff4841b          	addiw	s0,s1,-1
   104ba:	00044e63          	bltz	s0,104d6 <__call_exitprocs+0x56>
   104be:	048e                	slli	s1,s1,0x3
   104c0:	94ca                	add	s1,s1,s2
   104c2:	020b0663          	beqz	s6,104ee <__call_exitprocs+0x6e>
   104c6:	2084b783          	ld	a5,520(s1)
   104ca:	03678263          	beq	a5,s6,104ee <__call_exitprocs+0x6e>
   104ce:	347d                	addiw	s0,s0,-1
   104d0:	14e1                	addi	s1,s1,-8
   104d2:	ff3418e3          	bne	s0,s3,104c2 <__call_exitprocs+0x42>
   104d6:	60a6                	ld	ra,72(sp)
   104d8:	6406                	ld	s0,64(sp)
   104da:	74e2                	ld	s1,56(sp)
   104dc:	7942                	ld	s2,48(sp)
   104de:	79a2                	ld	s3,40(sp)
   104e0:	7a02                	ld	s4,32(sp)
   104e2:	6ae2                	ld	s5,24(sp)
   104e4:	6b42                	ld	s6,16(sp)
   104e6:	6ba2                	ld	s7,8(sp)
   104e8:	6c02                	ld	s8,0(sp)
   104ea:	6161                	addi	sp,sp,80
   104ec:	8082                	ret
   104ee:	00892783          	lw	a5,8(s2)
   104f2:	6498                	ld	a4,8(s1)
   104f4:	37fd                	addiw	a5,a5,-1
   104f6:	04878263          	beq	a5,s0,1053a <__call_exitprocs+0xba>
   104fa:	0004b423          	sd	zero,8(s1)
   104fe:	db61                	beqz	a4,104ce <__call_exitprocs+0x4e>
   10500:	31092783          	lw	a5,784(s2)
   10504:	008a16bb          	sllw	a3,s4,s0
   10508:	00892b83          	lw	s7,8(s2)
   1050c:	8ff5                	and	a5,a5,a3
   1050e:	2781                	sext.w	a5,a5
   10510:	eb99                	bnez	a5,10526 <__call_exitprocs+0xa6>
   10512:	9702                	jalr	a4
   10514:	00892783          	lw	a5,8(s2)
   10518:	f97799e3          	bne	a5,s7,104aa <__call_exitprocs+0x2a>
   1051c:	1f8c3783          	ld	a5,504(s8)
   10520:	fb2787e3          	beq	a5,s2,104ce <__call_exitprocs+0x4e>
   10524:	b759                	j	104aa <__call_exitprocs+0x2a>
   10526:	31492783          	lw	a5,788(s2)
   1052a:	1084b583          	ld	a1,264(s1)
   1052e:	8ff5                	and	a5,a5,a3
   10530:	2781                	sext.w	a5,a5
   10532:	e799                	bnez	a5,10540 <__call_exitprocs+0xc0>
   10534:	8556                	mv	a0,s5
   10536:	9702                	jalr	a4
   10538:	bff1                	j	10514 <__call_exitprocs+0x94>
   1053a:	00892423          	sw	s0,8(s2)
   1053e:	b7c1                	j	104fe <__call_exitprocs+0x7e>
   10540:	852e                	mv	a0,a1
   10542:	9702                	jalr	a4
   10544:	bfc1                	j	10514 <__call_exitprocs+0x94>

0000000000010546 <_exit>:
   10546:	4581                	li	a1,0
   10548:	4601                	li	a2,0
   1054a:	4681                	li	a3,0
   1054c:	4701                	li	a4,0
   1054e:	4781                	li	a5,0
   10550:	05d00893          	li	a7,93
   10554:	00000073          	ecall
   10558:	00054363          	bltz	a0,1055e <_exit+0x18>
   1055c:	a001                	j	1055c <_exit+0x16>
   1055e:	1141                	addi	sp,sp,-16
   10560:	e022                	sd	s0,0(sp)
   10562:	842a                	mv	s0,a0
   10564:	e406                	sd	ra,8(sp)
   10566:	4080043b          	negw	s0,s0
   1056a:	00000097          	auipc	ra,0x0
   1056e:	00c080e7          	jalr	12(ra) # 10576 <__errno>
   10572:	c100                	sw	s0,0(a0)
   10574:	a001                	j	10574 <_exit+0x2e>

0000000000010576 <__errno>:
   10576:	00001797          	auipc	a5,0x1
   1057a:	7aa78793          	addi	a5,a5,1962 # 11d20 <_impure_ptr>
   1057e:	6388                	ld	a0,0(a5)
   10580:	8082                	ret
