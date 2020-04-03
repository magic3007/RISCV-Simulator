
hello.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	17250513          	addi	a0,a0,370 # 1022c <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	12e30067          	jr	302(t1) # 101f0 <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	c4c18193          	addi	gp,gp,-948 # 11d18 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	bac50513          	addi	a0,a0,-1108 # 11c80 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	be460613          	addi	a2,a2,-1052 # 11cc0 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	1e6080e7          	jalr	486(ra) # 102ce <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	13c50513          	addi	a0,a0,316 # 1022c <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	0f8080e7          	jalr	248(ra) # 101f0 <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	164080e7          	jalr	356(ra) # 10264 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	07a080e7          	jalr	122(ra) # 10188 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	0ea30067          	jr	234(t1) # 10200 <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	b627c783          	lbu	a5,-1182(a5) # 11c80 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	3c450513          	addi	a0,a0,964 # 114f8 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	b2f70b23          	sb	a5,-1226(a4) # 11c80 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	b2f70423          	sb	a5,-1240(a4) # 11c80 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	b1a58593          	addi	a1,a1,-1254 # 11c88 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	38250513          	addi	a0,a0,898 # 114f8 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <main>:
int b = 2;
int result;

// result:  6

int main(){
   10188:	fe010113          	addi	sp,sp,-32
   1018c:	00813c23          	sd	s0,24(sp)
   10190:	02010413          	addi	s0,sp,32
	for(int i = 0; i < a; i++)
   10194:	fe042623          	sw	zero,-20(s0)
   10198:	0300006f          	j	101c8 <main+0x40>
		result += b;
   1019c:	000127b7          	lui	a5,0x12
   101a0:	cb87a703          	lw	a4,-840(a5) # 11cb8 <result>
   101a4:	000127b7          	lui	a5,0x12
   101a8:	c747a783          	lw	a5,-908(a5) # 11c74 <b>
   101ac:	00f707bb          	addw	a5,a4,a5
   101b0:	0007871b          	sext.w	a4,a5
   101b4:	000127b7          	lui	a5,0x12
   101b8:	cae7ac23          	sw	a4,-840(a5) # 11cb8 <result>
	for(int i = 0; i < a; i++)
   101bc:	fec42783          	lw	a5,-20(s0)
   101c0:	0017879b          	addiw	a5,a5,1
   101c4:	fef42623          	sw	a5,-20(s0)
   101c8:	000127b7          	lui	a5,0x12
   101cc:	c707a703          	lw	a4,-912(a5) # 11c70 <a>
   101d0:	fec42783          	lw	a5,-20(s0)
   101d4:	0007879b          	sext.w	a5,a5
   101d8:	fce7c2e3          	blt	a5,a4,1019c <main+0x14>
	return 0;
   101dc:	00000793          	li	a5,0
   101e0:	00078513          	mv	a0,a5
   101e4:	01813403          	ld	s0,24(sp)
   101e8:	02010113          	addi	sp,sp,32
   101ec:	00008067          	ret

00000000000101f0 <atexit>:
   101f0:	85aa                	mv	a1,a0
   101f2:	4681                	li	a3,0
   101f4:	4601                	li	a2,0
   101f6:	4501                	li	a0,0
   101f8:	00000317          	auipc	t1,0x0
   101fc:	18030067          	jr	384(t1) # 10378 <__register_exitproc>

0000000000010200 <exit>:
   10200:	1141                	addi	sp,sp,-16
   10202:	4581                	li	a1,0
   10204:	e022                	sd	s0,0(sp)
   10206:	e406                	sd	ra,8(sp)
   10208:	842a                	mv	s0,a0
   1020a:	00000097          	auipc	ra,0x0
   1020e:	1ea080e7          	jalr	490(ra) # 103f4 <__call_exitprocs>
   10212:	00002797          	auipc	a5,0x2
   10216:	a4e78793          	addi	a5,a5,-1458 # 11c60 <_global_impure_ptr>
   1021a:	6388                	ld	a0,0(a5)
   1021c:	6d3c                	ld	a5,88(a0)
   1021e:	c391                	beqz	a5,10222 <exit+0x22>
   10220:	9782                	jalr	a5
   10222:	8522                	mv	a0,s0
   10224:	00000097          	auipc	ra,0x0
   10228:	296080e7          	jalr	662(ra) # 104ba <_exit>

000000000001022c <__libc_fini_array>:
   1022c:	1101                	addi	sp,sp,-32
   1022e:	e822                	sd	s0,16(sp)
   10230:	00001797          	auipc	a5,0x1
   10234:	2e878793          	addi	a5,a5,744 # 11518 <__fini_array_end>
   10238:	00001417          	auipc	s0,0x1
   1023c:	2d840413          	addi	s0,s0,728 # 11510 <__init_array_end>
   10240:	8f81                	sub	a5,a5,s0
   10242:	e426                	sd	s1,8(sp)
   10244:	ec06                	sd	ra,24(sp)
   10246:	4037d493          	srai	s1,a5,0x3
   1024a:	c881                	beqz	s1,1025a <__libc_fini_array+0x2e>
   1024c:	17e1                	addi	a5,a5,-8
   1024e:	943e                	add	s0,s0,a5
   10250:	601c                	ld	a5,0(s0)
   10252:	14fd                	addi	s1,s1,-1
   10254:	1461                	addi	s0,s0,-8
   10256:	9782                	jalr	a5
   10258:	fce5                	bnez	s1,10250 <__libc_fini_array+0x24>
   1025a:	60e2                	ld	ra,24(sp)
   1025c:	6442                	ld	s0,16(sp)
   1025e:	64a2                	ld	s1,8(sp)
   10260:	6105                	addi	sp,sp,32
   10262:	8082                	ret

0000000000010264 <__libc_init_array>:
   10264:	1101                	addi	sp,sp,-32
   10266:	e822                	sd	s0,16(sp)
   10268:	e04a                	sd	s2,0(sp)
   1026a:	00001417          	auipc	s0,0x1
   1026e:	29240413          	addi	s0,s0,658 # 114fc <__preinit_array_end>
   10272:	00001917          	auipc	s2,0x1
   10276:	28a90913          	addi	s2,s2,650 # 114fc <__preinit_array_end>
   1027a:	40890933          	sub	s2,s2,s0
   1027e:	ec06                	sd	ra,24(sp)
   10280:	e426                	sd	s1,8(sp)
   10282:	40395913          	srai	s2,s2,0x3
   10286:	00090963          	beqz	s2,10298 <__libc_init_array+0x34>
   1028a:	4481                	li	s1,0
   1028c:	601c                	ld	a5,0(s0)
   1028e:	0485                	addi	s1,s1,1
   10290:	0421                	addi	s0,s0,8
   10292:	9782                	jalr	a5
   10294:	fe991ce3          	bne	s2,s1,1028c <__libc_init_array+0x28>
   10298:	00001417          	auipc	s0,0x1
   1029c:	26840413          	addi	s0,s0,616 # 11500 <__init_array_start>
   102a0:	00001917          	auipc	s2,0x1
   102a4:	27090913          	addi	s2,s2,624 # 11510 <__init_array_end>
   102a8:	40890933          	sub	s2,s2,s0
   102ac:	40395913          	srai	s2,s2,0x3
   102b0:	00090963          	beqz	s2,102c2 <__libc_init_array+0x5e>
   102b4:	4481                	li	s1,0
   102b6:	601c                	ld	a5,0(s0)
   102b8:	0485                	addi	s1,s1,1
   102ba:	0421                	addi	s0,s0,8
   102bc:	9782                	jalr	a5
   102be:	fe991ce3          	bne	s2,s1,102b6 <__libc_init_array+0x52>
   102c2:	60e2                	ld	ra,24(sp)
   102c4:	6442                	ld	s0,16(sp)
   102c6:	64a2                	ld	s1,8(sp)
   102c8:	6902                	ld	s2,0(sp)
   102ca:	6105                	addi	sp,sp,32
   102cc:	8082                	ret

00000000000102ce <memset>:
   102ce:	433d                	li	t1,15
   102d0:	872a                	mv	a4,a0
   102d2:	02c37163          	bgeu	t1,a2,102f4 <memset+0x26>
   102d6:	00f77793          	andi	a5,a4,15
   102da:	e3c1                	bnez	a5,1035a <memset+0x8c>
   102dc:	e1bd                	bnez	a1,10342 <memset+0x74>
   102de:	ff067693          	andi	a3,a2,-16
   102e2:	8a3d                	andi	a2,a2,15
   102e4:	96ba                	add	a3,a3,a4
   102e6:	e30c                	sd	a1,0(a4)
   102e8:	e70c                	sd	a1,8(a4)
   102ea:	0741                	addi	a4,a4,16
   102ec:	fed76de3          	bltu	a4,a3,102e6 <memset+0x18>
   102f0:	e211                	bnez	a2,102f4 <memset+0x26>
   102f2:	8082                	ret
   102f4:	40c306b3          	sub	a3,t1,a2
   102f8:	068a                	slli	a3,a3,0x2
   102fa:	00000297          	auipc	t0,0x0
   102fe:	9696                	add	a3,a3,t0
   10300:	00a68067          	jr	10(a3)
   10304:	00b70723          	sb	a1,14(a4)
   10308:	00b706a3          	sb	a1,13(a4)
   1030c:	00b70623          	sb	a1,12(a4)
   10310:	00b705a3          	sb	a1,11(a4)
   10314:	00b70523          	sb	a1,10(a4)
   10318:	00b704a3          	sb	a1,9(a4)
   1031c:	00b70423          	sb	a1,8(a4)
   10320:	00b703a3          	sb	a1,7(a4)
   10324:	00b70323          	sb	a1,6(a4)
   10328:	00b702a3          	sb	a1,5(a4)
   1032c:	00b70223          	sb	a1,4(a4)
   10330:	00b701a3          	sb	a1,3(a4)
   10334:	00b70123          	sb	a1,2(a4)
   10338:	00b700a3          	sb	a1,1(a4)
   1033c:	00b70023          	sb	a1,0(a4)
   10340:	8082                	ret
   10342:	0ff5f593          	andi	a1,a1,255
   10346:	00859693          	slli	a3,a1,0x8
   1034a:	8dd5                	or	a1,a1,a3
   1034c:	01059693          	slli	a3,a1,0x10
   10350:	8dd5                	or	a1,a1,a3
   10352:	02059693          	slli	a3,a1,0x20
   10356:	8dd5                	or	a1,a1,a3
   10358:	b759                	j	102de <memset+0x10>
   1035a:	00279693          	slli	a3,a5,0x2
   1035e:	00000297          	auipc	t0,0x0
   10362:	9696                	add	a3,a3,t0
   10364:	8286                	mv	t0,ra
   10366:	fa2680e7          	jalr	-94(a3)
   1036a:	8096                	mv	ra,t0
   1036c:	17c1                	addi	a5,a5,-16
   1036e:	8f1d                	sub	a4,a4,a5
   10370:	963e                	add	a2,a2,a5
   10372:	f8c371e3          	bgeu	t1,a2,102f4 <memset+0x26>
   10376:	b79d                	j	102dc <memset+0xe>

0000000000010378 <__register_exitproc>:
   10378:	00002797          	auipc	a5,0x2
   1037c:	8e878793          	addi	a5,a5,-1816 # 11c60 <_global_impure_ptr>
   10380:	6398                	ld	a4,0(a5)
   10382:	1f873783          	ld	a5,504(a4)
   10386:	c3b1                	beqz	a5,103ca <__register_exitproc+0x52>
   10388:	4798                	lw	a4,8(a5)
   1038a:	487d                	li	a6,31
   1038c:	06e84263          	blt	a6,a4,103f0 <__register_exitproc+0x78>
   10390:	c505                	beqz	a0,103b8 <__register_exitproc+0x40>
   10392:	00371813          	slli	a6,a4,0x3
   10396:	983e                	add	a6,a6,a5
   10398:	10c83823          	sd	a2,272(a6)
   1039c:	3107a883          	lw	a7,784(a5)
   103a0:	4605                	li	a2,1
   103a2:	00e6163b          	sllw	a2,a2,a4
   103a6:	00c8e8b3          	or	a7,a7,a2
   103aa:	3117a823          	sw	a7,784(a5)
   103ae:	20d83823          	sd	a3,528(a6)
   103b2:	4689                	li	a3,2
   103b4:	02d50063          	beq	a0,a3,103d4 <__register_exitproc+0x5c>
   103b8:	00270693          	addi	a3,a4,2
   103bc:	068e                	slli	a3,a3,0x3
   103be:	2705                	addiw	a4,a4,1
   103c0:	c798                	sw	a4,8(a5)
   103c2:	97b6                	add	a5,a5,a3
   103c4:	e38c                	sd	a1,0(a5)
   103c6:	4501                	li	a0,0
   103c8:	8082                	ret
   103ca:	20070793          	addi	a5,a4,512
   103ce:	1ef73c23          	sd	a5,504(a4)
   103d2:	bf5d                	j	10388 <__register_exitproc+0x10>
   103d4:	3147a683          	lw	a3,788(a5)
   103d8:	4501                	li	a0,0
   103da:	8e55                	or	a2,a2,a3
   103dc:	00270693          	addi	a3,a4,2
   103e0:	068e                	slli	a3,a3,0x3
   103e2:	2705                	addiw	a4,a4,1
   103e4:	30c7aa23          	sw	a2,788(a5)
   103e8:	c798                	sw	a4,8(a5)
   103ea:	97b6                	add	a5,a5,a3
   103ec:	e38c                	sd	a1,0(a5)
   103ee:	8082                	ret
   103f0:	557d                	li	a0,-1
   103f2:	8082                	ret

00000000000103f4 <__call_exitprocs>:
   103f4:	715d                	addi	sp,sp,-80
   103f6:	00002797          	auipc	a5,0x2
   103fa:	86a78793          	addi	a5,a5,-1942 # 11c60 <_global_impure_ptr>
   103fe:	e062                	sd	s8,0(sp)
   10400:	0007bc03          	ld	s8,0(a5)
   10404:	f44e                	sd	s3,40(sp)
   10406:	f052                	sd	s4,32(sp)
   10408:	ec56                	sd	s5,24(sp)
   1040a:	e85a                	sd	s6,16(sp)
   1040c:	e486                	sd	ra,72(sp)
   1040e:	e0a2                	sd	s0,64(sp)
   10410:	fc26                	sd	s1,56(sp)
   10412:	f84a                	sd	s2,48(sp)
   10414:	e45e                	sd	s7,8(sp)
   10416:	8aaa                	mv	s5,a0
   10418:	8b2e                	mv	s6,a1
   1041a:	4a05                	li	s4,1
   1041c:	59fd                	li	s3,-1
   1041e:	1f8c3903          	ld	s2,504(s8)
   10422:	02090463          	beqz	s2,1044a <__call_exitprocs+0x56>
   10426:	00892483          	lw	s1,8(s2)
   1042a:	fff4841b          	addiw	s0,s1,-1
   1042e:	00044e63          	bltz	s0,1044a <__call_exitprocs+0x56>
   10432:	048e                	slli	s1,s1,0x3
   10434:	94ca                	add	s1,s1,s2
   10436:	020b0663          	beqz	s6,10462 <__call_exitprocs+0x6e>
   1043a:	2084b783          	ld	a5,520(s1)
   1043e:	03678263          	beq	a5,s6,10462 <__call_exitprocs+0x6e>
   10442:	347d                	addiw	s0,s0,-1
   10444:	14e1                	addi	s1,s1,-8
   10446:	ff3418e3          	bne	s0,s3,10436 <__call_exitprocs+0x42>
   1044a:	60a6                	ld	ra,72(sp)
   1044c:	6406                	ld	s0,64(sp)
   1044e:	74e2                	ld	s1,56(sp)
   10450:	7942                	ld	s2,48(sp)
   10452:	79a2                	ld	s3,40(sp)
   10454:	7a02                	ld	s4,32(sp)
   10456:	6ae2                	ld	s5,24(sp)
   10458:	6b42                	ld	s6,16(sp)
   1045a:	6ba2                	ld	s7,8(sp)
   1045c:	6c02                	ld	s8,0(sp)
   1045e:	6161                	addi	sp,sp,80
   10460:	8082                	ret
   10462:	00892783          	lw	a5,8(s2)
   10466:	6498                	ld	a4,8(s1)
   10468:	37fd                	addiw	a5,a5,-1
   1046a:	04878263          	beq	a5,s0,104ae <__call_exitprocs+0xba>
   1046e:	0004b423          	sd	zero,8(s1)
   10472:	db61                	beqz	a4,10442 <__call_exitprocs+0x4e>
   10474:	31092783          	lw	a5,784(s2)
   10478:	008a16bb          	sllw	a3,s4,s0
   1047c:	00892b83          	lw	s7,8(s2)
   10480:	8ff5                	and	a5,a5,a3
   10482:	2781                	sext.w	a5,a5
   10484:	eb99                	bnez	a5,1049a <__call_exitprocs+0xa6>
   10486:	9702                	jalr	a4
   10488:	00892783          	lw	a5,8(s2)
   1048c:	f97799e3          	bne	a5,s7,1041e <__call_exitprocs+0x2a>
   10490:	1f8c3783          	ld	a5,504(s8)
   10494:	fb2787e3          	beq	a5,s2,10442 <__call_exitprocs+0x4e>
   10498:	b759                	j	1041e <__call_exitprocs+0x2a>
   1049a:	31492783          	lw	a5,788(s2)
   1049e:	1084b583          	ld	a1,264(s1)
   104a2:	8ff5                	and	a5,a5,a3
   104a4:	2781                	sext.w	a5,a5
   104a6:	e799                	bnez	a5,104b4 <__call_exitprocs+0xc0>
   104a8:	8556                	mv	a0,s5
   104aa:	9702                	jalr	a4
   104ac:	bff1                	j	10488 <__call_exitprocs+0x94>
   104ae:	00892423          	sw	s0,8(s2)
   104b2:	b7c1                	j	10472 <__call_exitprocs+0x7e>
   104b4:	852e                	mv	a0,a1
   104b6:	9702                	jalr	a4
   104b8:	bfc1                	j	10488 <__call_exitprocs+0x94>

00000000000104ba <_exit>:
   104ba:	4581                	li	a1,0
   104bc:	4601                	li	a2,0
   104be:	4681                	li	a3,0
   104c0:	4701                	li	a4,0
   104c2:	4781                	li	a5,0
   104c4:	05d00893          	li	a7,93
   104c8:	00000073          	ecall
   104cc:	00054363          	bltz	a0,104d2 <_exit+0x18>
   104d0:	a001                	j	104d0 <_exit+0x16>
   104d2:	1141                	addi	sp,sp,-16
   104d4:	e022                	sd	s0,0(sp)
   104d6:	842a                	mv	s0,a0
   104d8:	e406                	sd	ra,8(sp)
   104da:	4080043b          	negw	s0,s0
   104de:	00000097          	auipc	ra,0x0
   104e2:	00c080e7          	jalr	12(ra) # 104ea <__errno>
   104e6:	c100                	sw	s0,0(a0)
   104e8:	a001                	j	104e8 <_exit+0x2e>

00000000000104ea <__errno>:
   104ea:	00001797          	auipc	a5,0x1
   104ee:	78e78793          	addi	a5,a5,1934 # 11c78 <_impure_ptr>
   104f2:	6388                	ld	a0,0(a5)
   104f4:	8082                	ret
