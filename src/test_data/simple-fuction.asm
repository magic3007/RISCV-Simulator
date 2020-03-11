
simple-fuction.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	21650513          	addi	a0,a0,534 # 102d0 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	1d230067          	jr	466(t1) # 10294 <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	cec18193          	addi	gp,gp,-788 # 11db8 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	c6c50513          	addi	a0,a0,-916 # 11d40 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	c9c60613          	addi	a2,a2,-868 # 11d78 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	28a080e7          	jalr	650(ra) # 10372 <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	1e050513          	addi	a0,a0,480 # 102d0 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	19c080e7          	jalr	412(ra) # 10294 <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	208080e7          	jalr	520(ra) # 10308 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	156080e7          	jalr	342(ra) # 10264 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	18e30067          	jr	398(t1) # 102a4 <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	c227c783          	lbu	a5,-990(a5) # 11d40 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	46850513          	addi	a0,a0,1128 # 1159c <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	bef70b23          	sb	a5,-1034(a4) # 11d40 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	bef70423          	sb	a5,-1048(a4) # 11d40 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	bda58593          	addi	a1,a1,-1062 # 11d48 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	42650513          	addi	a0,a0,1062 # 1159c <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <fuction>:
int result[10]={1,2,3,4,5,6,7,8,9,10};

//result:11 12 13 14 15 1 2 3 4 5

void fuction()
{
   10188:	fe010113          	addi	sp,sp,-32
   1018c:	00813c23          	sd	s0,24(sp)
   10190:	02010413          	addi	s0,sp,32
	int i=0;
   10194:	fe042623          	sw	zero,-20(s0)
	for(i=0;i<5;i++)
   10198:	fe042623          	sw	zero,-20(s0)
   1019c:	0480006f          	j	101e4 <fuction+0x5c>
		result[i]=result[i]+10;
   101a0:	000117b7          	lui	a5,0x11
   101a4:	5b878713          	addi	a4,a5,1464 # 115b8 <__fini_array_end>
   101a8:	fec42783          	lw	a5,-20(s0)
   101ac:	00279793          	slli	a5,a5,0x2
   101b0:	00f707b3          	add	a5,a4,a5
   101b4:	0007a783          	lw	a5,0(a5)
   101b8:	00a7879b          	addiw	a5,a5,10
   101bc:	0007871b          	sext.w	a4,a5
   101c0:	000117b7          	lui	a5,0x11
   101c4:	5b878693          	addi	a3,a5,1464 # 115b8 <__fini_array_end>
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
   101f0:	fae7d8e3          	bge	a5,a4,101a0 <fuction+0x18>
	for(i=5;i<10;i++)
   101f4:	00500793          	li	a5,5
   101f8:	fef42623          	sw	a5,-20(s0)
   101fc:	0480006f          	j	10244 <fuction+0xbc>
		result[i]=result[i]-5;
   10200:	000117b7          	lui	a5,0x11
   10204:	5b878713          	addi	a4,a5,1464 # 115b8 <__fini_array_end>
   10208:	fec42783          	lw	a5,-20(s0)
   1020c:	00279793          	slli	a5,a5,0x2
   10210:	00f707b3          	add	a5,a4,a5
   10214:	0007a783          	lw	a5,0(a5)
   10218:	ffb7879b          	addiw	a5,a5,-5
   1021c:	0007871b          	sext.w	a4,a5
   10220:	000117b7          	lui	a5,0x11
   10224:	5b878693          	addi	a3,a5,1464 # 115b8 <__fini_array_end>
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
   10250:	fae7d8e3          	bge	a5,a4,10200 <fuction+0x78>
}
   10254:	00000013          	nop
   10258:	01813403          	ld	s0,24(sp)
   1025c:	02010113          	addi	sp,sp,32
   10260:	00008067          	ret

0000000000010264 <main>:

int main()
{
   10264:	ff010113          	addi	sp,sp,-16
   10268:	00113423          	sd	ra,8(sp)
   1026c:	00813023          	sd	s0,0(sp)
   10270:	01010413          	addi	s0,sp,16
	fuction();
   10274:	00000097          	auipc	ra,0x0
   10278:	f14080e7          	jalr	-236(ra) # 10188 <fuction>
	return 0;
   1027c:	00000793          	li	a5,0
   10280:	00078513          	mv	a0,a5
   10284:	00813083          	ld	ra,8(sp)
   10288:	00013403          	ld	s0,0(sp)
   1028c:	01010113          	addi	sp,sp,16
   10290:	00008067          	ret

0000000000010294 <atexit>:
   10294:	85aa                	mv	a1,a0
   10296:	4681                	li	a3,0
   10298:	4601                	li	a2,0
   1029a:	4501                	li	a0,0
   1029c:	00000317          	auipc	t1,0x0
   102a0:	18030067          	jr	384(t1) # 1041c <__register_exitproc>

00000000000102a4 <exit>:
   102a4:	1141                	addi	sp,sp,-16
   102a6:	4581                	li	a1,0
   102a8:	e022                	sd	s0,0(sp)
   102aa:	e406                	sd	ra,8(sp)
   102ac:	842a                	mv	s0,a0
   102ae:	00000097          	auipc	ra,0x0
   102b2:	1ea080e7          	jalr	490(ra) # 10498 <__call_exitprocs>
   102b6:	00002797          	auipc	a5,0x2
   102ba:	a7278793          	addi	a5,a5,-1422 # 11d28 <_global_impure_ptr>
   102be:	6388                	ld	a0,0(a5)
   102c0:	6d3c                	ld	a5,88(a0)
   102c2:	c391                	beqz	a5,102c6 <exit+0x22>
   102c4:	9782                	jalr	a5
   102c6:	8522                	mv	a0,s0
   102c8:	00000097          	auipc	ra,0x0
   102cc:	296080e7          	jalr	662(ra) # 1055e <_exit>

00000000000102d0 <__libc_fini_array>:
   102d0:	1101                	addi	sp,sp,-32
   102d2:	e822                	sd	s0,16(sp)
   102d4:	00001797          	auipc	a5,0x1
   102d8:	2e478793          	addi	a5,a5,740 # 115b8 <__fini_array_end>
   102dc:	00001417          	auipc	s0,0x1
   102e0:	2d440413          	addi	s0,s0,724 # 115b0 <__init_array_end>
   102e4:	8f81                	sub	a5,a5,s0
   102e6:	e426                	sd	s1,8(sp)
   102e8:	ec06                	sd	ra,24(sp)
   102ea:	4037d493          	srai	s1,a5,0x3
   102ee:	c881                	beqz	s1,102fe <__libc_fini_array+0x2e>
   102f0:	17e1                	addi	a5,a5,-8
   102f2:	943e                	add	s0,s0,a5
   102f4:	601c                	ld	a5,0(s0)
   102f6:	14fd                	addi	s1,s1,-1
   102f8:	1461                	addi	s0,s0,-8
   102fa:	9782                	jalr	a5
   102fc:	fce5                	bnez	s1,102f4 <__libc_fini_array+0x24>
   102fe:	60e2                	ld	ra,24(sp)
   10300:	6442                	ld	s0,16(sp)
   10302:	64a2                	ld	s1,8(sp)
   10304:	6105                	addi	sp,sp,32
   10306:	8082                	ret

0000000000010308 <__libc_init_array>:
   10308:	1101                	addi	sp,sp,-32
   1030a:	e822                	sd	s0,16(sp)
   1030c:	e04a                	sd	s2,0(sp)
   1030e:	00001417          	auipc	s0,0x1
   10312:	29240413          	addi	s0,s0,658 # 115a0 <__init_array_start>
   10316:	00001917          	auipc	s2,0x1
   1031a:	28a90913          	addi	s2,s2,650 # 115a0 <__init_array_start>
   1031e:	40890933          	sub	s2,s2,s0
   10322:	ec06                	sd	ra,24(sp)
   10324:	e426                	sd	s1,8(sp)
   10326:	40395913          	srai	s2,s2,0x3
   1032a:	00090963          	beqz	s2,1033c <__libc_init_array+0x34>
   1032e:	4481                	li	s1,0
   10330:	601c                	ld	a5,0(s0)
   10332:	0485                	addi	s1,s1,1
   10334:	0421                	addi	s0,s0,8
   10336:	9782                	jalr	a5
   10338:	fe991ce3          	bne	s2,s1,10330 <__libc_init_array+0x28>
   1033c:	00001417          	auipc	s0,0x1
   10340:	26440413          	addi	s0,s0,612 # 115a0 <__init_array_start>
   10344:	00001917          	auipc	s2,0x1
   10348:	26c90913          	addi	s2,s2,620 # 115b0 <__init_array_end>
   1034c:	40890933          	sub	s2,s2,s0
   10350:	40395913          	srai	s2,s2,0x3
   10354:	00090963          	beqz	s2,10366 <__libc_init_array+0x5e>
   10358:	4481                	li	s1,0
   1035a:	601c                	ld	a5,0(s0)
   1035c:	0485                	addi	s1,s1,1
   1035e:	0421                	addi	s0,s0,8
   10360:	9782                	jalr	a5
   10362:	fe991ce3          	bne	s2,s1,1035a <__libc_init_array+0x52>
   10366:	60e2                	ld	ra,24(sp)
   10368:	6442                	ld	s0,16(sp)
   1036a:	64a2                	ld	s1,8(sp)
   1036c:	6902                	ld	s2,0(sp)
   1036e:	6105                	addi	sp,sp,32
   10370:	8082                	ret

0000000000010372 <memset>:
   10372:	433d                	li	t1,15
   10374:	872a                	mv	a4,a0
   10376:	02c37163          	bgeu	t1,a2,10398 <memset+0x26>
   1037a:	00f77793          	andi	a5,a4,15
   1037e:	e3c1                	bnez	a5,103fe <memset+0x8c>
   10380:	e1bd                	bnez	a1,103e6 <memset+0x74>
   10382:	ff067693          	andi	a3,a2,-16
   10386:	8a3d                	andi	a2,a2,15
   10388:	96ba                	add	a3,a3,a4
   1038a:	e30c                	sd	a1,0(a4)
   1038c:	e70c                	sd	a1,8(a4)
   1038e:	0741                	addi	a4,a4,16
   10390:	fed76de3          	bltu	a4,a3,1038a <memset+0x18>
   10394:	e211                	bnez	a2,10398 <memset+0x26>
   10396:	8082                	ret
   10398:	40c306b3          	sub	a3,t1,a2
   1039c:	068a                	slli	a3,a3,0x2
   1039e:	00000297          	auipc	t0,0x0
   103a2:	9696                	add	a3,a3,t0
   103a4:	00a68067          	jr	10(a3)
   103a8:	00b70723          	sb	a1,14(a4)
   103ac:	00b706a3          	sb	a1,13(a4)
   103b0:	00b70623          	sb	a1,12(a4)
   103b4:	00b705a3          	sb	a1,11(a4)
   103b8:	00b70523          	sb	a1,10(a4)
   103bc:	00b704a3          	sb	a1,9(a4)
   103c0:	00b70423          	sb	a1,8(a4)
   103c4:	00b703a3          	sb	a1,7(a4)
   103c8:	00b70323          	sb	a1,6(a4)
   103cc:	00b702a3          	sb	a1,5(a4)
   103d0:	00b70223          	sb	a1,4(a4)
   103d4:	00b701a3          	sb	a1,3(a4)
   103d8:	00b70123          	sb	a1,2(a4)
   103dc:	00b700a3          	sb	a1,1(a4)
   103e0:	00b70023          	sb	a1,0(a4)
   103e4:	8082                	ret
   103e6:	0ff5f593          	andi	a1,a1,255
   103ea:	00859693          	slli	a3,a1,0x8
   103ee:	8dd5                	or	a1,a1,a3
   103f0:	01059693          	slli	a3,a1,0x10
   103f4:	8dd5                	or	a1,a1,a3
   103f6:	02059693          	slli	a3,a1,0x20
   103fa:	8dd5                	or	a1,a1,a3
   103fc:	b759                	j	10382 <memset+0x10>
   103fe:	00279693          	slli	a3,a5,0x2
   10402:	00000297          	auipc	t0,0x0
   10406:	9696                	add	a3,a3,t0
   10408:	8286                	mv	t0,ra
   1040a:	fa2680e7          	jalr	-94(a3)
   1040e:	8096                	mv	ra,t0
   10410:	17c1                	addi	a5,a5,-16
   10412:	8f1d                	sub	a4,a4,a5
   10414:	963e                	add	a2,a2,a5
   10416:	f8c371e3          	bgeu	t1,a2,10398 <memset+0x26>
   1041a:	b79d                	j	10380 <memset+0xe>

000000000001041c <__register_exitproc>:
   1041c:	00002797          	auipc	a5,0x2
   10420:	90c78793          	addi	a5,a5,-1780 # 11d28 <_global_impure_ptr>
   10424:	6398                	ld	a4,0(a5)
   10426:	1f873783          	ld	a5,504(a4)
   1042a:	c3b1                	beqz	a5,1046e <__register_exitproc+0x52>
   1042c:	4798                	lw	a4,8(a5)
   1042e:	487d                	li	a6,31
   10430:	06e84263          	blt	a6,a4,10494 <__register_exitproc+0x78>
   10434:	c505                	beqz	a0,1045c <__register_exitproc+0x40>
   10436:	00371813          	slli	a6,a4,0x3
   1043a:	983e                	add	a6,a6,a5
   1043c:	10c83823          	sd	a2,272(a6)
   10440:	3107a883          	lw	a7,784(a5)
   10444:	4605                	li	a2,1
   10446:	00e6163b          	sllw	a2,a2,a4
   1044a:	00c8e8b3          	or	a7,a7,a2
   1044e:	3117a823          	sw	a7,784(a5)
   10452:	20d83823          	sd	a3,528(a6)
   10456:	4689                	li	a3,2
   10458:	02d50063          	beq	a0,a3,10478 <__register_exitproc+0x5c>
   1045c:	00270693          	addi	a3,a4,2
   10460:	068e                	slli	a3,a3,0x3
   10462:	2705                	addiw	a4,a4,1
   10464:	c798                	sw	a4,8(a5)
   10466:	97b6                	add	a5,a5,a3
   10468:	e38c                	sd	a1,0(a5)
   1046a:	4501                	li	a0,0
   1046c:	8082                	ret
   1046e:	20070793          	addi	a5,a4,512
   10472:	1ef73c23          	sd	a5,504(a4)
   10476:	bf5d                	j	1042c <__register_exitproc+0x10>
   10478:	3147a683          	lw	a3,788(a5)
   1047c:	4501                	li	a0,0
   1047e:	8e55                	or	a2,a2,a3
   10480:	00270693          	addi	a3,a4,2
   10484:	068e                	slli	a3,a3,0x3
   10486:	2705                	addiw	a4,a4,1
   10488:	30c7aa23          	sw	a2,788(a5)
   1048c:	c798                	sw	a4,8(a5)
   1048e:	97b6                	add	a5,a5,a3
   10490:	e38c                	sd	a1,0(a5)
   10492:	8082                	ret
   10494:	557d                	li	a0,-1
   10496:	8082                	ret

0000000000010498 <__call_exitprocs>:
   10498:	715d                	addi	sp,sp,-80
   1049a:	00002797          	auipc	a5,0x2
   1049e:	88e78793          	addi	a5,a5,-1906 # 11d28 <_global_impure_ptr>
   104a2:	e062                	sd	s8,0(sp)
   104a4:	0007bc03          	ld	s8,0(a5)
   104a8:	f44e                	sd	s3,40(sp)
   104aa:	f052                	sd	s4,32(sp)
   104ac:	ec56                	sd	s5,24(sp)
   104ae:	e85a                	sd	s6,16(sp)
   104b0:	e486                	sd	ra,72(sp)
   104b2:	e0a2                	sd	s0,64(sp)
   104b4:	fc26                	sd	s1,56(sp)
   104b6:	f84a                	sd	s2,48(sp)
   104b8:	e45e                	sd	s7,8(sp)
   104ba:	8aaa                	mv	s5,a0
   104bc:	8b2e                	mv	s6,a1
   104be:	4a05                	li	s4,1
   104c0:	59fd                	li	s3,-1
   104c2:	1f8c3903          	ld	s2,504(s8)
   104c6:	02090463          	beqz	s2,104ee <__call_exitprocs+0x56>
   104ca:	00892483          	lw	s1,8(s2)
   104ce:	fff4841b          	addiw	s0,s1,-1
   104d2:	00044e63          	bltz	s0,104ee <__call_exitprocs+0x56>
   104d6:	048e                	slli	s1,s1,0x3
   104d8:	94ca                	add	s1,s1,s2
   104da:	020b0663          	beqz	s6,10506 <__call_exitprocs+0x6e>
   104de:	2084b783          	ld	a5,520(s1)
   104e2:	03678263          	beq	a5,s6,10506 <__call_exitprocs+0x6e>
   104e6:	347d                	addiw	s0,s0,-1
   104e8:	14e1                	addi	s1,s1,-8
   104ea:	ff3418e3          	bne	s0,s3,104da <__call_exitprocs+0x42>
   104ee:	60a6                	ld	ra,72(sp)
   104f0:	6406                	ld	s0,64(sp)
   104f2:	74e2                	ld	s1,56(sp)
   104f4:	7942                	ld	s2,48(sp)
   104f6:	79a2                	ld	s3,40(sp)
   104f8:	7a02                	ld	s4,32(sp)
   104fa:	6ae2                	ld	s5,24(sp)
   104fc:	6b42                	ld	s6,16(sp)
   104fe:	6ba2                	ld	s7,8(sp)
   10500:	6c02                	ld	s8,0(sp)
   10502:	6161                	addi	sp,sp,80
   10504:	8082                	ret
   10506:	00892783          	lw	a5,8(s2)
   1050a:	6498                	ld	a4,8(s1)
   1050c:	37fd                	addiw	a5,a5,-1
   1050e:	04878263          	beq	a5,s0,10552 <__call_exitprocs+0xba>
   10512:	0004b423          	sd	zero,8(s1)
   10516:	db61                	beqz	a4,104e6 <__call_exitprocs+0x4e>
   10518:	31092783          	lw	a5,784(s2)
   1051c:	008a16bb          	sllw	a3,s4,s0
   10520:	00892b83          	lw	s7,8(s2)
   10524:	8ff5                	and	a5,a5,a3
   10526:	2781                	sext.w	a5,a5
   10528:	eb99                	bnez	a5,1053e <__call_exitprocs+0xa6>
   1052a:	9702                	jalr	a4
   1052c:	00892783          	lw	a5,8(s2)
   10530:	f97799e3          	bne	a5,s7,104c2 <__call_exitprocs+0x2a>
   10534:	1f8c3783          	ld	a5,504(s8)
   10538:	fb2787e3          	beq	a5,s2,104e6 <__call_exitprocs+0x4e>
   1053c:	b759                	j	104c2 <__call_exitprocs+0x2a>
   1053e:	31492783          	lw	a5,788(s2)
   10542:	1084b583          	ld	a1,264(s1)
   10546:	8ff5                	and	a5,a5,a3
   10548:	2781                	sext.w	a5,a5
   1054a:	e799                	bnez	a5,10558 <__call_exitprocs+0xc0>
   1054c:	8556                	mv	a0,s5
   1054e:	9702                	jalr	a4
   10550:	bff1                	j	1052c <__call_exitprocs+0x94>
   10552:	00892423          	sw	s0,8(s2)
   10556:	b7c1                	j	10516 <__call_exitprocs+0x7e>
   10558:	852e                	mv	a0,a1
   1055a:	9702                	jalr	a4
   1055c:	bfc1                	j	1052c <__call_exitprocs+0x94>

000000000001055e <_exit>:
   1055e:	4581                	li	a1,0
   10560:	4601                	li	a2,0
   10562:	4681                	li	a3,0
   10564:	4701                	li	a4,0
   10566:	4781                	li	a5,0
   10568:	05d00893          	li	a7,93
   1056c:	00000073          	ecall
   10570:	00054363          	bltz	a0,10576 <_exit+0x18>
   10574:	a001                	j	10574 <_exit+0x16>
   10576:	1141                	addi	sp,sp,-16
   10578:	e022                	sd	s0,0(sp)
   1057a:	842a                	mv	s0,a0
   1057c:	e406                	sd	ra,8(sp)
   1057e:	4080043b          	negw	s0,s0
   10582:	00000097          	auipc	ra,0x0
   10586:	00c080e7          	jalr	12(ra) # 1058e <__errno>
   1058a:	c100                	sw	s0,0(a0)
   1058c:	a001                	j	1058c <_exit+0x2e>

000000000001058e <__errno>:
   1058e:	00001797          	auipc	a5,0x1
   10592:	7aa78793          	addi	a5,a5,1962 # 11d38 <_impure_ptr>
   10596:	6388                	ld	a0,0(a5)
   10598:	8082                	ret
